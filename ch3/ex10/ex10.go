/*
Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.
*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("412123456789"))
}

// comma inserts commas in a non-negative decimal integer string
func comma(s string) string {
	len := len(s)
	remain := len % 3
	if remain == 0 {
		remain = 3
	}
	var buf bytes.Buffer
	buf.WriteString(s[:remain])
	for i := remain; i < len; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
