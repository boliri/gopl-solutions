package main

import (
	"fmt"
	"os"

	"reverse"
)

func main() {
	var s string = os.Args[1]

	fmt.Printf("Original string: %v\n", s)

	b := []byte(s)
	reverse.Reverse([]byte(b))
	fmt.Printf("Reversed sequence: %v\n\n", string(b))
}
