/*
Enhance comma so that it deals correctly with floating-point numbers and an optional sign.
*/

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-123456789.243535634"))
}

func comma(s string) string {
	var buf bytes.Buffer
	// Determine whether there is a sign, and offset the starting index by 1 if so
	start := 0
	if sign := s[0]; sign == '-' || sign == '+' {
		buf.WriteByte(sign)
		start++
	}
	// Find the length of the integer part
	size := strings.Index(s, ".")
	if size == -1 {
		size = len(s)
	}
	// Determine number of integer-part digits before the first comma, offsetting any sign
	remain := (size - start) % 3
	if remain == 0 {
		remain = 3
	}
	remain += start
	// Write the integer part to the buffer
	buf.WriteString(s[start:remain])
	for i := remain; i < size; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	// Write the fractional part to the buffer
	buf.WriteString(s[size:])
	return buf.String()
}
