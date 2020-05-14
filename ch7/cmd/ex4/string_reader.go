// The strings.NewReader function returns a value that satisfies the io.Reader interface (and
// others) by reading from its argument, a string. Implement a simple version of NewReader yourself,
// and use it to make the HTML parser (§5.2) take input from a string.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	HTML := `<html lang="en">
	<head></head>
	<body></body>
	</html>`
	reader := &StringReader{str: HTML}
	stringReaderRoot, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	bufferRoot, _ := html.Parse(strings.NewReader(HTML))
	fmt.Printf("Roots equal? %t", stringReaderRoot.Data == bufferRoot.Data)
}

type StringReader struct {
	str      string
	readFrom int
}

func (s *StringReader) Read(p []byte) (n int, err error) {
	if s.readFrom >= len(s.str) {
		return 0, io.EOF
	}
	n = copy(p, s.str[s.readFrom:])
	s.readFrom += n
	return n, nil
}
