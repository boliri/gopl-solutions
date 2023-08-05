package popcount

import (
	"testing"
)

const v = 500000


func BenchmarkPopCountSingleExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountSingleExpression(v)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearRightmostNonZeroBit(v)
	}
}