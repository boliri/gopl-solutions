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

	fmt.Printf("New set created: %s\n", &s)
	fmt.Printf("Set length: %d\n\n", s.Len())

	fmt.Printf("Removing elements from %s...\n", &s)
	s.Remove(2)
	s.Remove(3)
	fmt.Printf("Elements removed. Current state: %s\n\n", &s)

	fmt.Printf("Copying %s to a new set...\n", &s)
	t := s.Copy()
	fmt.Println("Set copied\n")

	fmt.Printf("Clearing set %s...\n", &s)
	s.Clear()
	fmt.Println("Set cleared\n")

	fmt.Printf("Set S -> %s\n", &s)
	fmt.Printf("Set T -> %s\n\n", t)
}