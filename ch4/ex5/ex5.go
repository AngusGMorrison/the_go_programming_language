/*
Write an in-place function to eliminate duplicates in a []string slice.
*/

package main

import "fmt"

func main() {
	strings := []string{
		"hello",
		"hi",
		"hello",
		"howdy",
		"hello",
	}
	fmt.Println(dedup(strings))
}

func dedup(strings []string) []string {
	i := 0
	seen := make(map[string]bool)
	for _, s := range strings {
		if !seen[s] {
			seen[s] = true
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
