// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

package main

import (
	"fmt"
	"os"

	"replace"
)

func main() {
	in := os.Args[1]

	var fn string
	if len(os.Args) > 2 {
		fn = os.Args[2]
	}

	var f replace.MutateFn
	switch fn {
	case "--caesar":
		f = replace.Caesar
	case "--hex":
		f = replace.Hex
	default:
		f = replace.Noop
	}

	out := replace.Expand(in, f)
	fmt.Println(out)
}

//!-
