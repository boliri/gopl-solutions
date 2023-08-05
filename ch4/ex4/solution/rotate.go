// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package rotate

//!+
// rotates a slice of ints left by N positions
func LeftRotate(s []int, n int) {
	t := make([]int, len(s))
	copy(t, s)

	copy(s[len(s)-n:], t[:n])
	copy(s[:n+1], t[n:])
}
//!-