package palindrome

import (
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	lastIdx := s.Len() / 2

	isEven := s.Len()%2 == 0
	if !isEven {
		lastIdx -= 1
	}

	for i := 0; i <= lastIdx; i++ {
		j := s.Len() - 1 - i // like index i, but moving from the end of the slice to the beginning

		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}

	return true
}
