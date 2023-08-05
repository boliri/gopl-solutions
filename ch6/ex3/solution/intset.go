// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith returns a new set with elements existing in s and t.
func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	var x IntSet

	for i, tword := range t.words {
		if i < len(s.words) {
			x.words = append(x.words, s.words[i] & tword)
		} else {
			break
		}
	}

	return &x
}

// DifferenceWith returns a new set with elements existing in s but not in t.
func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	var x IntSet

	for i, sword := range s.words {
		if i < len(t.words) {
			x.words = append(x.words, sword &^ t.words[i])
		} else {
			x.words = append(x.words, sword)
		}
	}

	return &x
}

// SymmetricDifference returns a new set with elements existing in s and t, but not in both.
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	var x IntSet

	for i, sword := range s.words {
		if i < len(t.words) {
			x.words = append(x.words, sword ^ t.words[i])
		} else {
			x.words = append(x.words, sword)
		}
	}

	return &x
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
