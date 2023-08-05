// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"gopl.io/ch5/links"
)

type urlT struct {
	link  string
	depth int
}

var maxDepth *int = flag.Int("depth", 3, "max traversal's depth per URL passed as argument")

// !+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url urlT) []urlT {
	fmt.Println(url.link)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url.link)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}

	var extracted []urlT
	for _, l := range list {
		extracted = append(extracted, urlT{link: l, depth: url.depth + 1})
	}
	return extracted
}

//!-sema

// !+
func main() {
	flag.Parse()

	worklist := make(chan urlT)
	var wg sync.WaitGroup

	// Start with the command-line arguments, and make sure to skip the depth flag if it exists
	initialUrls := make(map[string]bool)
	for _, in := range os.Args[1:] {
		if strings.HasPrefix(in, "-depth") {
			continue
		}
		initialUrls[in] = true
	}

	// we need to increment the waitgroup so the worklist channel is not closed prematurely in one of the upcoming goroutines
	wg.Add(len(initialUrls))

	go func() {
		for url := range initialUrls {
			worklist <- urlT{link: url, depth: 0}
		}
	}()

	// Close the worklist channel when there are no more urls to fetch
	go func() {
		wg.Wait()
		close(worklist)
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for url := range worklist {
		if !seen[url.link] {
			seen[url.link] = true

			// do not increment the waitgroup if the url is among the initial ones
			// otherwise the worklist channel will never get closed and we'll fall into a deadlock
			if _, ok := initialUrls[url.link]; !ok {
				wg.Add(1)
			}

			go func(url urlT) {
				for _, u := range crawl(url) {
					if u.depth > *maxDepth {
						continue
					}
					worklist <- u
				}
				wg.Done()
			}(url)
		}
	}
}

//!-
