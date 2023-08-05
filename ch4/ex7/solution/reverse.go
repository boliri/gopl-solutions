// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package reverse

import (
	"unicode/utf8"
)

//!+
// reverses a slice of bytes that represents a UTF-8-encoded string in place
func Reverse(s []byte) {
	t := s[:]

	for len(t) > 0 {
		_, size := utf8.DecodeRune(t)

		for size > 0 {
			// Move the last byte of the rune to the end of slice s as much positions as slice t allows
			for i := size - 1; i < (len(t) - 1); i++ {
				s[i], s[i+1] = s[i+1], s[i]
			}
			size--
			t = t[:len(t)-1]
		}
	}
}