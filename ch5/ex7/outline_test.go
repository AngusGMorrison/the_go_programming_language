package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// Test that the output of forEachNode is valid HTML than can be successfully parsed
func TestForEachNode(t *testing.T) {
	testHTML := `<html lang="en">
	  <head>
		<meta charset="utf-8"/>
		<title>Go!</title>
	  </head>
	  <body>
	    <!-- A comment -->
	  	<p class="text">Text node</p>
	  </body>
	</html>`

	doc, _ := html.Parse(strings.NewReader(testHTML))
	out = new(bytes.Buffer)
	forEachNode(doc, startElement, endElement)
	_, err := html.Parse(out.(*bytes.Buffer))
	if err != nil {
		t.Errorf("outline output couldn't be parsed: %v", err)
	}
}
