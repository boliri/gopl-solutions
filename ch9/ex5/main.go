package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var bounces uint64

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go send(ch1, ch2)
	go send(ch2, ch1)
	ch1 <- struct{}{}

	ticker := time.Tick(time.Second)
	for {
		<-ticker
		fmt.Printf("%d bounces in the last second\n", bounces)
		atomic.AddUint64(&bounces, -atomic.LoadUint64(&bounces))
	}
}

func send(out chan<- struct{}, in <-chan struct{}) {
	for v := range in {
		out <- v
		atomic.AddUint64(&bounces, 1)
	}
}
