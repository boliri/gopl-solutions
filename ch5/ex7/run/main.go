// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

package main

import (
	"fmt"
	"os"

	"prettyhtml"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "prettyhtml: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(prettyhtml.Prettify(doc))
}

//!-
