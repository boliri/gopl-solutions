// Package intset provides ways to manipulate sets of integers through multiple data structures.
package intset

// IntSet defines the set of methods a data structure must implement to be considered as such
type IntSet[T MapSet | BitVector] interface {
	Add(x int)
	DifferenceWith(t *T) *T
	Elems() []uint64
	Has(x int) bool
	IntersectWith(t *T) *T
	String() string
	SymmetricDifference(t *T) *T
	UnionWith(t *T)
}
