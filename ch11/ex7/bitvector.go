package intset

import (
	"bytes"
	"fmt"
)

// A BitVector is a set of small non-negative integers built upon a bit vector.
// Its zero value represents the empty set.
//
// Satisfies the IntSet interface.
type BitVector struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *BitVector) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *BitVector) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *BitVector) UnionWith(t *BitVector) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith returns a new set with elements existing in s and t.
func (s *BitVector) IntersectWith(t *BitVector) *BitVector {
	var x BitVector

	for i, tword := range t.words {
		if i < len(s.words) {
			x.words = append(x.words, s.words[i]&tword)
		} else {
			break
		}
	}

	return &x
}

// DifferenceWith returns a new set with elements existing in s but not in t.
func (s *BitVector) DifferenceWith(t *BitVector) *BitVector {
	var x BitVector

	for i, sword := range s.words {
		if i < len(t.words) {
			x.words = append(x.words, sword&^t.words[i])
		} else {
			x.words = append(x.words, sword)
		}
	}

	return &x
}

// SymmetricDifference returns a new set with elements existing in s and t, but not in both.
func (s *BitVector) SymmetricDifference(t *BitVector) *BitVector {
	var x BitVector

	for i, sword := range s.words {
		if i < len(t.words) {
			x.words = append(x.words, sword^t.words[i])
		} else {
			x.words = append(x.words, sword)
		}
	}

	return &x
}

// Elems returns a slice of all the elements in the set.
func (s *BitVector) Elems() []uint64 {
	var elems []uint64

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				elems = append(elems, uint64(i*64+j))
			}
		}
	}

	return elems
}

// String returns the set as a string of the form "{1 2 3}".
func (s *BitVector) String() string {
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
