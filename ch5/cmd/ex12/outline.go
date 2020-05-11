// The startElement and endElement functions in gopl.io/ch5/outline2 (§5.5) share a global variable,
// depth. Turn them into anonymous functions that share a variable local to the outline function.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url\n", os.Args[0])
		os.Exit(1)
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "getting %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "%s responded %d\n", os.Args[1], resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	outline(doc)
}

func outline(root *html.Node) {
	var depth int
	var startElement func(n *html.Node)
	var endElement func(n *html.Node)

	startElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	endElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	forEachElement(root, startElement, endElement)
}

func forEachElement(root *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(root)
	}

	for c := root.FirstChild; c != nil; c = c.NextSibling {
		forEachElement(c, pre, post)
	}

	if post != nil {
		post(root)
	}
}
