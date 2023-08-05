package intset

import (
	"fmt"
	"sort"
	"strings"
)

// A MapSet is a set of small non-negative integers built upon a Go map.
// Its zero value represents the empty set.
//
// Satisfies the IntSet interface.
type MapSet map[uint64]bool

// Has reports whether the set contains the non-negative value x.
func (s *MapSet) Has(x int) bool {
	return (*s)[uint64(x)]
}

// Add adds the non-negative value x to the set.
func (s *MapSet) Add(x int) {
	(*s)[uint64(x)] = true
}

// UnionWith sets s to the union of s and t.
func (s *MapSet) UnionWith(t *MapSet) {
	for elem := range *t {
		(*s)[uint64(elem)] = true
	}
}

// IntersectWith returns a new set with elements existing in s and t.
func (s *MapSet) IntersectWith(t *MapSet) *MapSet {
	x := MapSet{}

	for elem := range *s {
		if (*t)[elem] {
			x[elem] = true
		}
	}

	return &x
}

// DifferenceWith returns a new set with elements existing in s but not in t.
func (s *MapSet) DifferenceWith(t *MapSet) *MapSet {
	x := MapSet{}

	for elem := range *s {
		if !(*t)[elem] {
			x[elem] = true
		}
	}

	return &x
}

// SymmetricDifference returns a new set with elements existing in s and t, but not in both.
func (s *MapSet) SymmetricDifference(t *MapSet) *MapSet {
	x := MapSet{}

	for elem := range *s {
		if !(*t)[elem] {
			x[elem] = true
		}
	}

	for elem := range *t {
		if !(*s)[elem] {
			x[elem] = true
		}
	}

	return &x
}

// Elems returns a slice of all the elements in the set.
func (s *MapSet) Elems() []uint64 {
	var elems []uint64
	for elem := range *s {
		elems = append(elems, elem)
	}

	sort.Slice(elems, func(i, j int) bool { return elems[i] < elems[j] })
	return elems
}

// String returns the set as a string of the form "{1 2 3}".
func (s *MapSet) String() string {
	var buf strings.Builder
	buf.WriteString("{")
	for i, elem := range s.Elems() {
		buf.WriteString(fmt.Sprintf("%d", elem))
		if i < len(*s)-1 {
			buf.WriteString(" ")
		}
	}
	buf.WriteString("}")
	return buf.String()
}
