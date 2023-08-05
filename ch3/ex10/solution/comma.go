// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package comma

import "bytes"

//!+
// comma inserts commas in a non-negative decimal integer string.
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var r bytes.Buffer
	for i := 0; i <= n-1; i++ {
		r.WriteByte(s[i])

		addComma := len(s[i+1:]) % 3 == 0
		hasMore := i != n-1
		if addComma && hasMore {
			r.WriteRune(',')
		}
	}

	return r.String()
}

//!-
