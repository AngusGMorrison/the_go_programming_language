package charcount

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

const arrsize = utf8.UTFMax + 1

// Count reads from input and returns a count of each rune, the number of times a rune of len n
// bytes occurs, the number of invalid UTF-8 chracters and any error.
func Count(input io.Reader) (counts map[rune]int, utflens [arrsize]int, invalid int, err error) {
	counts = make(map[rune]int) // counts of Unicode characters
	in := bufio.NewReader(input)
	for {
		r, n, e := in.ReadRune() // returns rune, nbytes, error
		if e != nil {
			if e == io.EOF {
				break
			}
			err = e
			return
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflens[n]++
	}
	return
}
