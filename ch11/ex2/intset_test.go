package intset

import "testing"

func equal(s IntSet[BitVector], t IntSet[MapSet]) bool {
	tElems := t.Elems()
	for i, elem := range s.Elems() {
		if tElems[i] != elem {
			return false
		}
	}
	return true
}

func TestIntSet_MultipleImplementationsYieldSameResults(t *testing.T) {
	bv := &BitVector{}
	ms := &MapSet{}

	in := 10
	bv.Add(in)
	ms.Add(in)
	if !equal(bv, ms) {
		t.Logf("mismatch on Add(%d)", in)
		t.Logf("BitVector: %s", bv)
		t.Logf("MapSet:    %s", ms)
		t.Fail()
	}

	bv2 := &BitVector{}
	bv2.Add(100)
	ms2 := &MapSet{}
	ms2.Add(100)
	bv.UnionWith(bv2)
	ms.UnionWith(ms2)
	if !equal(bv, ms) {
		t.Logf("mismatch on %s ∪ %s\n", bv, bv2)
		t.Logf("BitVector: %s", bv)
		t.Logf("MapSet:    %s", ms)
		t.Fail()
	}

	diffBv := bv.DifferenceWith(bv2)
	diffMs := ms.DifferenceWith(ms2)
	if !equal(diffBv, diffMs) {
		t.Errorf("mismatch on %s - %s", bv, bv2)
		t.Errorf("BitVector: %v", diffBv)
		t.Errorf("MapSet:    %v", diffMs)
		t.Fail()
	}

	bv2 = &BitVector{}
	bv2.Add(10)
	ms2 = &MapSet{}
	ms2.Add(10)
	intersectBv := bv.IntersectWith(bv2)
	intersectMs := ms.IntersectWith(ms2)
	if !equal(intersectBv, intersectMs) {
		t.Errorf("mismatch on %s ∩ %s", bv, bv2)
		t.Errorf("BitVector: %v", intersectBv)
		t.Errorf("MapSet:    %v", intersectMs)
		t.Fail()
	}

	bv.Add(100)
	bv2 = &BitVector{}
	bv2.Add(10)
	bv2.Add(200)

	ms.Add(100)
	ms2 = &MapSet{}
	ms2.Add(10)
	ms2.Add(200)

	symmDiffBv := bv.SymmetricDifference(bv2)
	symmDiffMs := ms.SymmetricDifference(ms2)
	if !equal(symmDiffBv, symmDiffMs) {
		t.Errorf("mismatch on %s △ %s", bv, bv2)
		t.Errorf("BitVector: %v", intersectBv)
		t.Errorf("MapSet:    %v", intersectMs)
		t.Fail()
	}
}
