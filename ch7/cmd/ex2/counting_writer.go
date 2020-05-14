// Write a function CountingWriter with the signature below that, given an io.Writer, returns a new
// Writer that wraps the original, and a pointer to an int64 variable that at any moment contains
// the number of bytes written to the new Writer.
//
//		func CountingWriter(w io.Writer) (io.Writer, *int64)
//
package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	cw, count := CountingWriter(ioutil.Discard)
	cw.Write([]byte("hello"))
	fmt.Println(*count)

}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountWriter{w: w}
	return &cw, &cw.count
}

type CountWriter struct {
	count int64
	w     io.Writer
}

func (c *CountWriter) Write(p []byte) (int, error) {
	n, _ := c.w.Write(p)
	c.count += int64(n)
	return n, nil
}
