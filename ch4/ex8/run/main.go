// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

package main

import (
	"bufio"
	"os"

	"charcount"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	charcount.Count(in)
}

//!-
