package main

import (
	"fmt"
	"os"

	"squashws"
)

func main() {
	var s string = os.Args[1]

	fmt.Printf("Original string: %v\n", s)

	r := squashws.SquashWhitespaces([]byte(s))
	fmt.Printf("Whitespaces squashed: %s\n\n", r)
}
