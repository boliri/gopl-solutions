// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package squashws

import (
	"unicode"
)

//!+
// removes adjacent whitespaces in a UTF-8-encoded slice of bytes
func SquashWhitespaces(s []byte) []byte {
	var i, removed int

	for i < (len(s) - 1) {
		if i == (len(s) - 1 - removed) { break }

		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			s[i] = byte(' ')
			copy(s[i:], s[i+1:])
			removed++
		} else {
			i++
		}
	}

	return s[:len(s)-removed]
}
//!-