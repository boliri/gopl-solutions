// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package nodupes

//!+
// removes adjacent duplicates in a slice of strings
func RemoveAdjacentDupes(s []string) []string {
	if len(s) == 1 { return s }

	var i, removed int

	for i < (len(s) - 1) {
		if i == (len(s) - 1 - removed) { break }

		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			removed++
		} else {
			i++
		}
	}

	return s[:len(s)-removed]
}
//!-