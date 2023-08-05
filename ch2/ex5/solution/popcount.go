// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// This version of popcount works exactly the same than the original one. The only difference is that the table with pre-computed values
// is not used; instead, the bits sequence is ANDed with the same sequence minus 1, which clears the rightmost non-zero bit of the original sequence.
//!+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountSingleExpression(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountClearRightmostNonZeroBit(x uint64) int {
	var pop int

	for i:=0; i < 64; i++ {
		n := byte(x >> i)
		if (n - 1) == (n & (n-1)) {
			pop++
		}
	}
	
	return pop
}

//!-
