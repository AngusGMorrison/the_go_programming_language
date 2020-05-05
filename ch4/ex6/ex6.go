/*
Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace)
in a UTF-8-encoded []byte slice into a single ASCII space.
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	test := "asd   dfsdfsd  werwfsdfsd       dfsdf"
	squashed := squash([]byte(test))
	fmt.Println(string(squashed))
}

func squash(bytes []byte) []byte {
	for i := 0; i < len(bytes); i++ {
		spaces := 0
		for unicode.IsSpace(rune(bytes[i+spaces])) {
			spaces++
		}
		if spaces > 0 {
			bytes[i] = ' '
			bytes = append(bytes[:i+1], bytes[i+spaces:]...)
		}
	}
	return bytes
}
