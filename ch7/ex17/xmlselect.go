// Extend xmlselect so that elements may be selected not just by name, but by their attributes
// too, in the manner of CSS, so that, for instance, an element like <div id="page" class="wide">
// could be selected by a matching id or class as well as its name.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				names := mapXML(stack, func(x xml.StartElement) string { return x.Name.Local })
				fmt.Printf("%s: %s\n", strings.Join(names, " "), tok)
			}
		}
	}
}

func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if hasMatchingAttribute(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func hasMatchingAttribute(x xml.StartElement, y string) bool {
	if x.Name.Local == y {
		return true
	}
	for _, attr := range x.Attr {
		if attr.Value == y {
			return true
		}
	}
	return false
}

func mapXML(x []xml.StartElement, f func(xml.StartElement) string) []string {
	output := make([]string, len(x))
	for i, item := range x {
		output[i] = f(item)
	}
	return output
}
