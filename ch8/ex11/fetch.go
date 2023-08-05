package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

var done = make(chan struct{})

// requestCtx controls cancellations of all HTTP requests made through the fetch method
type requestCtx int

func (ctx *requestCtx) Deadline() (deadline time.Time, ok bool) { return }
func (ctx *requestCtx) Done() <-chan struct{}                   { return done }
func (ctx *requestCtx) Value(key any) any                       { return nil }

func (ctx *requestCtx) Err() error {
	select {
	case <-ctx.Done():
		return context.Canceled
	default:
		return nil
	}
}

var ctx *requestCtx

// !+
// fetch downloads the URL to local disk, pushes the url to ch and cancels any other fetch
// goroutines running if the push succeeds
func fetch(url string, ch chan<- string) {
	buf := &bytes.Reader{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// another goroutine might have won the race while the request was in flight - in that
		// case, a context.Canceled error is printed
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// if multiple, simultaneous fetch goroutines manage to succeed on making the HTTP GET request,
	// only one must win
	//
	// this select statement makes sure there's only one winner
	select {
	case done <- struct{}{}: // no winner yet
		close(done)
		defer func() { ch <- url }()
	case <-done: // someone else won the race, so let's bail out
		return
	}

	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	if err != nil {
		fmt.Println(err)
		return
	}
}

//!-

func main() {
	ch := make(chan string, 1)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	// wait for any request to resolve
	<-done

	// at this point all requests but one should've been cancelled, but we still need to wait for
	// the running goroutine to do its job
	winner := <-ch

	fmt.Printf("\n%s won the race!\n\n", winner)
}
