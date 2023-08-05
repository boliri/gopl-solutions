// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

package main

import (
	"fmt"
	"log"
	"os"

	"xkcd"
)

func main() {
	term := os.Args[1]
	comics, err := xkcd.GetComicsByTerm(term)

	if err != nil {
		log.Fatal(err)
	}

	for _, c := range comics {
		fmt.Printf("%s\n\n%s\n\n", c.GetURL(), c.Transcript)
		fmt.Print("---------------------\n\n")
	}
}

//!-
