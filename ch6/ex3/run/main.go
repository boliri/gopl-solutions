package main

import (
	"fmt"

	"intset"
)

func main() {
	fmt.Println("Creating two sets...")

	var s intset.IntSet
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)

	var t intset.IntSet
	t.Add(4)
	t.Add(5)
	t.Add(6)
	t.Add(7)

	fmt.Printf("Set S created: %s\n", &s)
	fmt.Printf("Set T created: %s\n\n", &t)

	fmt.Printf("Intersection of S and T   -> %s\n", s.IntersectWith(&t))
	fmt.Printf("Difference of S with T    -> %s\n", s.DifferenceWith(&t))
	fmt.Printf("Symmetric diff of S and T -> %s\n\n", s.SymmetricDifference(&t))
}