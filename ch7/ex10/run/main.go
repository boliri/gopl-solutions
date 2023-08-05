package main

import (
	"fmt"

	"palindrome"
)

type sequence []byte

func (s sequence) Len() int           { return len(s) }
func (s sequence) Less(i, j int) bool { return s[i] < s[j] }
func (s sequence) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	var s sequence

	s = sequence([]byte("ABBA"))
	fmt.Printf("Sequence: %s\n", s)
	fmt.Printf("Is a palindrome? -> %v\n\n", palindrome.IsPalindrome(&s))

	s = sequence([]byte("ALA"))
	fmt.Printf("Sequence: %s\n", s)
	fmt.Printf("Is a palindrome? -> %v\n\n", palindrome.IsPalindrome(&s))

	s = sequence([]byte("Hey"))
	fmt.Printf("Sequence: %s\n", s)
	fmt.Printf("Is a palindrome? -> %v\n\n", palindrome.IsPalindrome(&s))

	s = sequence([]byte{2, 3, 3, 2})
	fmt.Printf("Sequence: %v\n", s)
	fmt.Printf("Is a palindrome? -> %v\n\n", palindrome.IsPalindrome(&s))

	s = sequence([]byte{1, 2, 3, 4})
	fmt.Printf("Sequence: %v\n", s)
	fmt.Printf("Is a palindrome? -> %v\n\n", palindrome.IsPalindrome(&s))
}
