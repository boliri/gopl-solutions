package main

import (
	"fmt"

	"intset"
)

func main() {
	fmt.Printf("Creating new set of uint%d words...\n", intset.ARCH_BITS)

	var s intset.IntSet
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)

	fmt.Printf("New set created -> %s\n", &s)
}