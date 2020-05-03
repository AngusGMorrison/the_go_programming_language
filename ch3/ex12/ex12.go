/*
Write a function that reports whether two strings are anagrams of each other, that is, they contain
the same letters in a different order.
*/

package main

import "fmt"

func main() {
	fmt.Println(isAnagram("babcock♛", "kcaob♛cb"))
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var total rune
	for _, r := range s {
		total += r
	}
	for _, r := range t {
		total -= r
	}
	if total == 0 {
		return true
	}
	return false
}
