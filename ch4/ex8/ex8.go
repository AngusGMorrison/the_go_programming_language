/*
Modify charcount to count letters, digits and so on in their Unicode categories, using functions
like unicode.IsLetter.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	runeCounts := make(map[rune]int)   // Counts of Unicode characters
	letterCounts := make(map[rune]int) // Counts of letters
	digitCounts := make(map[rune]int)  // Counts of digits
	var utflen [utf8.UTFMax + 1]int    // Count of lengths of UTF-8 encodings
	invalid := 0                       // Count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // Returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("charCount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letterCounts[r]++
		} else if unicode.IsDigit(r) {
			digitCounts[r]++
		}
		runeCounts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for k, v := range runeCounts {
		fmt.Printf("%q\t%d\n", k, v)
	}
	fmt.Print("\nletter\tcount\n")
	for k, v := range letterCounts {
		fmt.Printf("%q\t%d\n", k, v)
	}
	fmt.Print("\ndigit\tcount\n")
	for k, v := range digitCounts {
		fmt.Printf("%q\t%d\n", k, v)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
