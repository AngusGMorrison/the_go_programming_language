// The sort.Interface type can be adopted to other uses. Write a function
// IsPalindrome(s sort.Interface) bool that reports whether the sequence s is a palindrome, in
// other words, reversing the sequence would not change it. Assume that the elements at indices
// i and j are equal if !s.Less(i, j) && !s.Less(j, i).
package palindrome

import "sort"

func IsPalindrome(s sort.Interface) bool {
	len := s.Len()
	for i := 0; i < len/2; i++ {
		j := len - 1 - i
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
