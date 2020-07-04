// TestRandomPalindromes only tests palindromes. Write a randomized test that generates a verifies
// non-palindromes.
package word

import (
	"math/rand"
	"testing"
	"time"
)

// Initialize a pseudo-random number generator.
func numberGenerator(t *testing.T) *rand.Rand {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	return rand.New(rand.NewSource(seed))
}

func randomRune(rng *rand.Rand) rune {
	return rune(rng.Intn(0x1000)) // random rune up to '\u0999'
}

// randomPalindrom returns a palindrome whose length and contents are derived from the
// pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := randomRune(rng)
		runes[i] = r
		runes[n-1-i] = r
	}
	punctdRunes := punctuate(runes, rng)
	return string(punctdRunes)
}

// punctuate pseudo-randomly inserts punctuation and spaces into a valid palindrome.
func punctuate(runes []rune, rng *rand.Rand) []rune {
	punct := []rune(`!@#%&*()-_][{}\;:/?.,§ `)
	punctdRunes := make([]rune, 0, len(runes))
	for i := 0; i < len(runes); {
		if rng.Intn(2)%2 == 0 {
			punctdRunes = append(punctdRunes, punct[rng.Intn(len(punct))])
		} else {
			punctdRunes = append(punctdRunes, runes[i])
			i++
		}
	}
	return punctdRunes
}

func TestRandomPalindromes(t *testing.T) {
	rng := numberGenerator(t)
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
