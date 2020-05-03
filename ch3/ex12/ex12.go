/*
Write a function that reports whether two strings are anagrams of each other, that is, they contain
the same letters in a different order.
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isAnagram("babcock♛", "Abbcck♛o"))
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var total rune
	for _, r := range s {
		total += unicode.ToLower(r)
	}
	for _, r := range t {
		total -= unicode.ToLower(r)
	}
	return total == 0
}
