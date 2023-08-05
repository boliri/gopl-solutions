// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// This program calculates how many bits are different in two SHA-256 hashes
//
// Example:
//  Inputs: hello world
//  Hashes (hex):
//    hello -> 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
//    world -> 486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7
//  Output: 112
//
//  Inputs: hello hello
//  Hashes (hex):
//    hello -> 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
//    hello -> 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
//  Output: 0
//
package bitsdiffsha256

import "math/bits"

//!+
func GetBitsDiff(h1, h2 [32]byte) int {
	var d int

	for i := 0; i < len(h1); i++ {
		var bseq1, bseq2 byte = h1[i], h2[i]
		diff := bseq1 ^ bseq2
		d += bits.OnesCount(uint(diff))
	}

	return d
}

//!-