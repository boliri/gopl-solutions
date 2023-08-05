// This program constructs a pipeline with as many channels as the system's resources allows, and
// sends a dummy value to traverse it entirely.
//
// I ran it in my own machine but eventually got bored of waiting for the system to run out of
// memory, so I just stopped the program. Using a Docker container with memory limits
// might have make things easier.
package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

var goros uint64
var start = time.Now()
var val = struct{}{}

func main() {
	out := make(chan struct{})
	in := make(chan struct{})

	go pipeline(out, in)
	in <- val
	for {
		// wait until the system runs out of resources, or until the user stops the program
	}
}

func pipeline(out, in chan struct{}) {
	// if the system runs out of resources, recover from the panic and print a summary
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(
				"%v took %.06fs to traverse a pipeline of %d channels\n\n",
				val, time.Since(start).Seconds(), goros,
			)
			os.Exit(0)
		}
	}()

	atomic.AddUint64(&goros, 1)

	// create the next stage
	nout := make(chan struct{})
	nin := out
	go pipeline(nout, nin)
	nin <- <-in

	for {
		// this infinite loop just prevents the garbage collector from reclaiming any resources
		// associated to this goroutine
	}
}
