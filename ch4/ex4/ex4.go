/*
Write a version of rotate that operates in a single pass.
*/

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(rotateLeft(s, 2))
}

// rotate rotates a slice left by pos positions, returning a new slice
func rotateLeft(s []int, pos int) []int {
	return append(s[pos:], s[:pos]...)
}
