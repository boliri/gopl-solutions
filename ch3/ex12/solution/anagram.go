// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// This program determines whether two strings are anagrams of each other or not
//
// Example:
//  Inputs: abc cba
//  Output: both strings are anagrams
//
//  Inputs: abc abca
//  Output: both strings are not anagrams
//
//  Inputs: abc abcd
//  Output: both strings are not anagrams
//
package anagram

import "strings"

//!+
func AreAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, c := range s1 {
		if strings.Count(s1, string(c)) != strings.Count(s2, string(c)) {
			return false
		}
	}

	return true
}

//!-