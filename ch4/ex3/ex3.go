/*
Rewrite reverse to use an array pointer instead of a slice.
*/

package main

import "fmt"

func main() {
	s := [5]int{1, 2, 3, 4, 5}
	reverse(&s)
	fmt.Println(s)
}

func reverse(s *[5]int) {
	for i, j := 0, 4; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
