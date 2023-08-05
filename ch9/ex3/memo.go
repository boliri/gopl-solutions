package memo

import (
	"context"
	"errors"
	"sync"
)

// Func is the type of the function to memoize.
type Func func(string, chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// !+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) GetCache() map[string]*entry {
	memo.mu.Lock()
	defer memo.mu.Unlock()
	return memo.cache
}

func (memo *Memo) Get(key string, done chan struct{}) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key, done)

		// CAVEAT: this check is tightly coupled to context-based cancellations
		if errors.Is(e.res.err, context.Canceled) {
			// although deleting a non-existing element from a map is a no-op, this sync via mutex
			// ensures that all goroutines calling Get() can see the same version of memo.cache, thus
			// avoiding removing the key from the copy of memo.cache stored in their CPU cache for
			// literally nothing
			memo.mu.Lock()
			delete(memo.cache, key)
			memo.mu.Unlock()

			return nil, e.res.err
		}

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}

//!-
