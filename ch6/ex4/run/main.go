package main

import (
	"fmt"

	"intset"
)

func main() {
	fmt.Println("Creating new set...")

	var s intset.IntSet
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)

	fmt.Printf("New set created -> %s\n", &s)

	fmt.Printf("Elements in set as slice -> %v\n\n", s.Elems())
}