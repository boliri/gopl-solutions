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

	fmt.Printf("New set created: %s\n\n", &s)

	nums := []int{100, 200, 300}
	fmt.Printf("Adding %v to set %s using AddAll method...\n", nums, &s)
	s.AddAll(nums...)
	fmt.Printf("Elements added. New set: %s\n\n", &s)
}