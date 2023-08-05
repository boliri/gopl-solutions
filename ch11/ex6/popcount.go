package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountTable returns the population count (number of set bits) of x using a pre-computed table.
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountShift returns the population count (number of set bits) of x by shifting the bits
// sequence 64 times to observe the rightmost bit.
func PopCountShift(x uint64) int {
	var pop int

	for i := 0; i < 64; i++ {
		pop += int(byte(x>>i)) % 2
	}

	return pop
}

// PopCountClearRightmostNonZeroBit returns the population count (number of set bits) of x by
// AND-ing the bits sequence with the same sequence minus 1, which clears the rightmost
// non-zero bit of the original sequence.
func PopCountClearRightmostNonZeroBit(x uint64) int {
	var pop int

	for i := 0; i < 64; i++ {
		n := byte(x >> i)
		if (n - 1) == (n & (n - 1)) {
			pop++
		}
	}

	return pop
}
