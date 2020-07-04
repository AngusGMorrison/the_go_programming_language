// Package word provides utilities for word games.
package word

import (
	"unicode"
)

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		// Allow non-letter characters but ignore punctuation and whitespace.
		if !unicode.IsPunct(r) && !unicode.IsSpace(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
