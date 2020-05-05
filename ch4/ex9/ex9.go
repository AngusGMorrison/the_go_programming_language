/*
Write a program wordfreq to report the frequency of each word in an input text file. Call
input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of
lines.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./ex9 <file/to/read>")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	wordFreq := countWords(file)
	fmt.Printf("%-16s %s\n", "word", "count")
	for k, v := range wordFreq {
		fmt.Printf("%-16s %d\n", k, v)
	}
}

func countWords(file *os.File) map[string]int {
	wordFreq := make(map[string]int)
	in := bufio.NewScanner(file)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		wordFreq[in.Text()]++
	}
	if err := in.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "ex9: %s", err)
	}
	return wordFreq
}
