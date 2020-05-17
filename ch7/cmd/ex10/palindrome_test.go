package palindrome

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	palins := map[*[]int]bool{
		{1, 2, 3, 2, 1}:    true,
		{1, 2, 3, 3, 2, 1}: true,
		{}:                 true,
		{1, 5, 2, 7, 5}:    false,
	}
	for slice, expected := range palins {
		isPalin := IsPalindrome(sort.IntSlice(*slice))
		if isPalin != expected {
			t.Errorf("%v: expected %t, got %t", *slice, expected, isPalin)
		}
	}
}
