// The LimitReader function in the io package accepts an io.Reader r and a number of bytes n, and
// returns another Reader that reads from r but reports an end-of-file condition after n bytes.
// Implement it.
//
//		func LimitReader(r io.Reader, n int64) io.Reader
//
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	reader := LimitReader(strings.NewReader("Hello"), 3)
	out := make([]byte, 5)
	_, err := reader.Read(out)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}

type LimReader struct {
	reader    io.Reader
	remaining int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimReader{
		reader:    r,
		remaining: n,
	}
}

func (l *LimReader) Read(p []byte) (n int, err error) {
	length := int64(len(p))
	if l.remaining < length {
		length = l.remaining
	}
	n, err = l.reader.Read(p[:length])
	if l.remaining -= int64(n); l.remaining == 0 {
		return n, io.EOF
	}
	return n, err
}
