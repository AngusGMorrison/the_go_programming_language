/*
Write an in-place function to eliminate duplicates in a []string slice.
*/

package main

import "fmt"

func main() {
	strings := []string{
		"hello",
		"hello",
		"hello",
		"hi",
		"hi",
		"hello",
		"howdy",
		"hello",
		"hello",
	}
	fmt.Println(removeAdjacentDups(strings))
}

func removeAdjacentDups(s []string) []string {
	// Iterate over all characters in the slice except the last one
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			s = append(s[:i], s[i+1:]...)
		} else {
			i++
		}
	}
	return s
}
