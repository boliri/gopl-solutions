// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

package main

import (
	"fmt"
	"os"
	"strings"

	"elemfinder"

	"golang.org/x/net/html"
)

func main() {
	in, elemId := os.Args[1], os.Args[2]

	inreader := strings.NewReader(in)
	doc, err := html.Parse(inreader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elemfinder: %v\n", err)
		os.Exit(1)
	}

	found := elemfinder.Find(doc, elemId)
	if found {
		fmt.Printf("found element with id=\"%s\" in html document\n\n", elemId)
	} else {
		fmt.Printf("couldn't find element with id=\"%s\" in html document\n\n", elemId)
	}
}

//!-
