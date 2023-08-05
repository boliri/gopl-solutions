package main

import (
	"fmt"
	"os"

	"anagram"
)

func main() {
	s1, s2 := os.Args[1], os.Args[2]
	areAnagrams := anagram.AreAnagrams(s1, s2)

	if areAnagrams {
		fmt.Printf("%s and %s are anagrams\n\n", s1, s2)
	} else {
		fmt.Printf("%s and %s are not anagrams\n\n", s1, s2)
	}
}
