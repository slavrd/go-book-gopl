package palindrome

import "sort"

func IsPalindrome(s sort.Interface) bool {

	l := s.Len()

	// check if sequence has at least 2 elements
	if l < 2 {
		return false
	}

	// check elements
	for i := 0; i <= l/2-1; i++ {
		if s.Less(i, l-1-i) || s.Less(l-1-i, i) {
			return false
		}
	}

	return true
}
