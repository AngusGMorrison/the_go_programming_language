// Write a function to print the contents of all text nodes in an HTML document tree.
// Do not descend into <script> or <stle> elements, since their contents are not visible in a web
// browser.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "print_text_nodes: %v\n", err)
		os.Exit(1)
	}
	printTextLinks(doc)
}

// printTextLinks visits each tree node recursively and prints its contents if it is a visible text
// node.
func printTextLinks(n *html.Node) {
	if n == nil || nodeNotVisible(n) {
		return
	}

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	printTextLinks(n.FirstChild)
	printTextLinks(n.NextSibling)
}

// nodeNotVisible returns false if the node won't be rendered in the browser
func nodeNotVisible(n *html.Node) bool {
	return n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style")
}
