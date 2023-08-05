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

import (
	"fmt"
	"strconv"
	"strings"
)

//!+
// comma inserts commas in positive / negative decimal integer / float string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func Comma(s string) string {
	var res string

	var sign string
	if _, err := strconv.ParseUint(string(s[0]), 10, 8); err != nil {
		sign = s[0:1]
		s = s[1:]
	}

	seqs := strings.Split(s, ".")
	res = fmt.Sprintf("%s%s", sign, comma(seqs[0]))
	if len(seqs) > 1 {
		res = fmt.Sprintf("%s.%s", res, seqs[1])
	}

	return res
}

//!-