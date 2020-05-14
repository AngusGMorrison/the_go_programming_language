// Using the ideas from ByteCounter, implement counters for words and for lines. You will find
// bufio.ScanWords useful.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	var w WordCounter
	count, _ := w.Write([]byte("hello"))
	fmt.Printf("%d words\n", count)
	fmt.Fprintf(&w, "hello, %s", "Angus")
	fmt.Printf("%d words\n", w)
	fmt.Println()

	var l LineCounter
	lines := `This is
	a three line
	string.`
	count, _ = l.Write([]byte(lines))
	fmt.Printf("%d lines\n", count)
	fmt.Fprint(&l, lines)
	fmt.Printf("%d lines\n", l)
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (wordCount int, err error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordCount++
	}
	*w += WordCounter(wordCount)
	return wordCount, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (lineCount int, err error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	for scanner.Scan() {
		lineCount++
	}
	*l += LineCounter(lineCount)
	return lineCount, nil
}
