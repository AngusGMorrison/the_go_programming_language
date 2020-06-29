// TestRandomPalindromes only tests palindromes. Write a randomized test that generates a verifies
// non-palindromes.
package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
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
	return string(runes)
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

func randomNonPunct(rng *rand.Rand) rune {
	r := randomRune(rng)
	for unicode.IsPunct(r) || unicode.IsSpace(r) {
		r = randomRune(rng)
	}
	return r
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25-2) + 2 // avoid single-rune words, which are always palindromes
	runes := make([]rune, n)
	for i := 0; i < n+1/2; i++ {
		// Avoid spaces and punct, which are ignored by IsPalindrome, potentially creating an
		// unintentional palindrome.
		first := randomNonPunct(rng)
		second := randomNonPunct(rng)
		for second == first {
			second = randomNonPunct(rng) // ensure opposing positions are always different
		}
		runes[i] = first
		runes[n-1-i] = second
	}
	return string(runes)
}

func TestRandomNonPalindromes(t *testing.T) {
	rng := numberGenerator(t)
	for i := 0; i < 1000; i++ {
		np := randomNonPalindrome(rng)
		if IsPalindrome(np) {
			t.Errorf("IsPalindrome(%q) = false", np)
		}
	}
}
