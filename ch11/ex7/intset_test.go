package intset

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	log.Printf("Random seed: %d", seed)

	os.Exit(m.Run())
}

// randomSetElems returns a slice of as much random ints as specified by size.
// All the slice's elements are guaranteed to be in the closed interval [min, max].
func randomSetElems(size, min, max int) []int {
	s := make([]int, 0, size)
	dupes := make(map[int]bool)
	for len(s) < size {
		n := rand.Intn((max - min + 1) + min)
		if _, ok := dupes[n]; ok {
			continue
		}

		s = append(s, n)
		dupes[n] = true
	}

	return s
}

var benchConfs = []struct {
	size, min, max int
}{
	{5, 0, 10},
	{50, 10, 100},
	{500, 100, 1000},
	{5000, 1000, 10000},
	{50000, 10000, 100000},
	{500000, 100000, 1000000},
	// These confs are too stressful for my machine, so I will leave them out for now
	// {5000000, 1000000, 10000000},
	// {50000000, 10000000, 100000000},
	// {500000000, 100000000, 1000000000},
	// {5000000000, 1000000000, 10000000000},
	// {50000000000, 10000000000, 100000000000},
	// {500000000000, 100000000000, 1000000000000},
}

func BenchmarkAdd(b *testing.B) {
	for _, conf := range benchConfs {
		bv := BitVector{}
		ms := MapSet{}
		elems := randomSetElems(conf.size, conf.min, conf.max)

		b.Run(
			fmt.Sprintf("type=%T_size=%d_min=%d_max=%d", bv, conf.size, conf.min, conf.max),
			func(b *testing.B) {
				for _, e := range elems {
					bv.Add(e)
				}
			},
		)
		b.Run(
			fmt.Sprintf("type=%T_size=%d_min=%d_max=%d", ms, conf.size, conf.min, conf.max),
			func(b *testing.B) {
				for _, e := range elems {
					ms.Add(e)
				}
			},
		)
	}
}

func BenchmarkUnionWith(b *testing.B) {
	for _, conf := range benchConfs {
		bv := BitVector{}
		ms := MapSet{}
		elems := randomSetElems(conf.size, conf.min, conf.max)
		for _, e := range elems {
			bv.Add(e)
			ms.Add(e)
		}

		elems2 := randomSetElems(conf.size, conf.min, conf.max)
		bv2 := BitVector{}
		ms2 := MapSet{}
		for _, e := range elems2 {
			bv2.Add(e)
			ms2.Add(e)
		}

		b.Run(
			fmt.Sprintf("type=%T_size=%d_min=%d_max=%d", bv, conf.size, conf.min, conf.max),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					bv.UnionWith(&bv2)
				}
			},
		)
		b.Run(
			fmt.Sprintf("type=%T_size=%d_min=%d_max=%d", ms, conf.size, conf.min, conf.max),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ms.UnionWith(&ms2)
				}
			},
		)
	}
}

func BenchmarkIntersectWith(b *testing.B) {
	for _, conf := range benchConfs {
		bv := BitVector{}
		ms := MapSet{}
		elems := randomSetElems(conf.size, conf.min, conf.max)
		for _, e := range elems {
			bv.Add(e)
			ms.Add(e)
		}

		elems2 := randomSetElems(conf.size, conf.min, conf.max)
		bv2 := BitVector{}
		ms2 := MapSet{}
		for _, e := range elems2 {
			bv2.Add(e)
			ms2.Add(e)
		}

		b.Run(
			fmt.Sprintf("type=%T_size=%d_min=%d_max=%d", bv, conf.size, conf.min, conf.max),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					bv.IntersectWith(&bv2)
				}
			},
		)
		b.Run(
			fmt.Sprintf("type=%T_size=%d_min=%d_max=%d", ms, conf.size, conf.min, conf.max),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ms.IntersectWith(&ms2)
				}
			},
		)
	}
}

func BenchmarkDifferenceWith(b *testing.B) {
	for _, conf := range benchConfs {
		bv := BitVector{}
		ms := MapSet{}
		elems := randomSetElems(conf.size, conf.min, conf.max)
		for _, e := range elems {
			bv.Add(e)
			ms.Add(e)
		}

		elems2 := randomSetElems(conf.size, conf.min, conf.max)
		bv2 := BitVector{}
		ms2 := MapSet{}
		for _, e := range elems2 {
			bv2.Add(e)
			ms2.Add(e)
		}

		b.Run(
			fmt.Sprintf("type=%T_size=%d_min=%d_max=%d", bv, conf.size, conf.min, conf.max),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					bv.DifferenceWith(&bv2)
				}
			},
		)
		b.Run(
			fmt.Sprintf("type=%T_size=%d_min=%d_max=%d", ms, conf.size, conf.min, conf.max),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ms.DifferenceWith(&ms2)
				}
			},
		)
	}
}

func BenchmarkSymmetricDifference(b *testing.B) {
	for _, conf := range benchConfs {
		bv := BitVector{}
		ms := MapSet{}
		elems := randomSetElems(conf.size, conf.min, conf.max)
		for _, e := range elems {
			bv.Add(e)
			ms.Add(e)
		}

		elems2 := randomSetElems(conf.size, conf.min, conf.max)
		bv2 := BitVector{}
		ms2 := MapSet{}
		for _, e := range elems2 {
			bv2.Add(e)
			ms2.Add(e)
		}

		b.Run(
			fmt.Sprintf("type=%T_size=%d_min=%d_max=%d", bv, conf.size, conf.min, conf.max),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					bv.SymmetricDifference(&bv2)
				}
			},
		)
		b.Run(
			fmt.Sprintf("type=%T_size=%d_min=%d_max=%d", ms, conf.size, conf.min, conf.max),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ms.SymmetricDifference(&ms2)
				}
			},
		)
	}
}
