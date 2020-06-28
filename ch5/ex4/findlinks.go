// Extend the visit functions so that it extracts other kinds of links from the document, such as
// images, scripts and style sheets.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	var match string
	if n.Type == html.ElementNode {
		if n.Data == "a" || n.Data == "link" {
			match = getMatchingAttr(n.Attr, "href")
		} else if n.Data == "img" || n.Data == "script" {
			match = getMatchingAttr(n.Attr, "src")
		}
	}

	if match != "" {
		links = append(links, match)
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

// getMatchingAttr searches a slice of Attribute for a key matching the key argument and returns
// its value.
func getMatchingAttr(attrs []html.Attribute, key string) string {
	for _, attr := range attrs {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
