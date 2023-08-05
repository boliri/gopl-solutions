package main

import (
	"fmt"
	"os"

	"nodupes"
)

func main() {
	var slice []string = os.Args[1:]

	fmt.Printf("Original strings sequence: %v\n", slice)

	r := nodupes.RemoveAdjacentDupes(slice)
	fmt.Printf("Adjacent dupes removed: %v\n\n", r)
}
