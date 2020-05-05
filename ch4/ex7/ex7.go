/*
Modify reverse to reverse the characters of a []byte slice that represents a UTF-8 encoded string,
in place. Can you do it without allocating new memory?
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	reversed := reverseUTF8([]byte("⚀⚁⚂⚃⚄⚅"))
	fmt.Println(string(reversed))
}

// Reverse each rune, then reverse the whole slice
func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverseBytes(b[i : i+size])
		i += size
	}
	reverseBytes(b)
	return b
}

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
