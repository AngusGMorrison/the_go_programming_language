// Write a function to populate a mapping from element names--p, div, span and so on--to the number
// of elements with that name in an HTML document.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "count_elements: %v\n", err)
		os.Exit(1)
	}
	counts := map[string]int{}
	countElements(doc, counts)
	for k, v := range counts {
		fmt.Printf("%-10s: %d\n", k, v)
	}
}

// count populates an external map with the count of each element type encountered.
func countElements(n *html.Node, counts map[string]int) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	countElements(n.FirstChild, counts)
	countElements(n.NextSibling, counts)
}
