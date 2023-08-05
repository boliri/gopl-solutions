package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"

	"countelements"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countelements: %v\n", err)
		os.Exit(1)
	}

	counts := make(map[string]int)
	countelements.Count(counts, doc)

	for tag, n := range counts {
		fmt.Printf("%s\t%d time(s)\n", tag, n)
	}
}
