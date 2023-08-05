// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

package main

import (
	"os"

	"wordfreq"
)

func main() {
	filename := os.Args[1]
	wordfreq.Count(filename)
}

//!-
