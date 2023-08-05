package memo_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	"memo"
)

var urls = []string{
	"https://golang.org",
	"https://godoc.org",
	"https://play.golang.org",
	"http://gopl.io",
	"https://golang.org",
	"https://godoc.org",
	"https://play.golang.org",
	"http://gopl.io",
}

// requestCtx controls cancellations of all HTTP requests made through the Extract method
type requestCtx struct {
	done <-chan struct{}
}

func (ctx *requestCtx) setDone(done <-chan struct{})            { ctx.done = done }
func (ctx *requestCtx) Deadline() (deadline time.Time, ok bool) { return }
func (ctx *requestCtx) Done() <-chan struct{}                   { return ctx.done }
func (ctx *requestCtx) Value(key any) any                       { return nil }

func (ctx *requestCtx) Err() error {
	select {
	case <-ctx.Done():
		return context.Canceled
	default:
		return nil
	}
}

func httpGetBody(url string, done chan struct{}) (interface{}, error) {
	ctx := &requestCtx{}
	ctx.setDone(done)

	buf := &bytes.Reader{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, buf)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, ch chan struct{}) (interface{}, error)
}

func Sequential(t *testing.T, m M, done chan struct{}) {
	//!+seq
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, done)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(t *testing.T, m M, done chan struct{}) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	var done chan struct{}
	Sequential(t, m, done)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	done := make(chan struct{})
	Concurrent(t, m, done)
}

func TestConcurrentWithCancellation(t *testing.T) {
	ticker := time.NewTicker(1 * time.Nanosecond)

	m := memo.New(httpGetBody)
	done := make(chan struct{})
	go Concurrent(t, m, done)

	// cancel as much requests as possible after the first tick
	<-ticker.C
	close(done)
	ticker.Stop()

	uniqueUrls := make(map[string]bool)
	for _, u := range urls {
		if _, ok := uniqueUrls[u]; ok {
			continue
		}
		uniqueUrls[u] = true
	}

	// the ticker is set to tick as earlier as Go permits (after 1ns) so we cache some requests,
	// but not all; however, Go can be blazingly fast so this test might not succeed consistently
	//
	// if you don't care about seeing the cache empty all the time, you can remove any references
	// to the ticker, and let the done channel be closed right after calling Concurrent()
	c := m.GetCache()
	if len(c) == len(uniqueUrls) {
		t.Errorf("Cache has %d elements, want %d or less", len(c), len(uniqueUrls)-1)
	} else {
		t.Errorf("Cache has %d elements after cancelling %d out of %d requests", len(c), len(uniqueUrls)-len(c), len(uniqueUrls))
	}
}
