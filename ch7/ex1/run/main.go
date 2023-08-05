package main

import (
	"counters"
	"fmt"
)

func main() {
	var wc counters.WordCounter
	var lc counters.LineCounter

	s := "Hi there!\nHow are you?"

	wc.Write([]byte(s))
	lc.Write([]byte(s))

	fmt.Printf("\"%s\"\n\n", s)
	fmt.Printf("%d words | %d lines\n\n", wc, lc)
}
