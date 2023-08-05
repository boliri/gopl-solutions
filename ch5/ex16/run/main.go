package main

import (
	"fmt"
	"os"

	"joiner"
)

func main() {
	sep := os.Args[1]
	strings := os.Args[2:]

	fmt.Println(joiner.Join(sep, strings...))
}