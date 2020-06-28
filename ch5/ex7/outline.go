// Develop startElement and endElement into a general HTML pretty-printer. Print comment nodes,
// text nodes, and the attributes of each element (<a href='...'>). Use short forms like <img/>
// instead of <img></img> when an element has no children. Write a test to ensure that the output
// can be parsed successfully. (See Chapter 11.)
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout // modified during testing

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url", os.Args[0])
		os.Exit(1)
	}

	if err := outline(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v", err)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s responded %d", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %v", url, err)
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

var depth int

func startElement(n *html.Node) {
	indent := depth * 2

	switch n.Type {
	case html.ElementNode:
		fmt.Fprintf(out, "%*s<%s %s", indent, "", n.Data, attrsToString(n.Attr))
		if n.FirstChild == nil {
			fmt.Fprint(out, "/>\n")
		} else {
			fmt.Fprint(out, ">\n")
			depth++
		}
	case html.CommentNode:
		fmt.Fprintf(out, "%*s<!--%s-->\n", indent, "", n.Data)
	case html.TextNode:
		if n.Data != "\n" {
			fmt.Fprintf(out, "%*s%s\n", indent, "", n.Data)
		}
	}
}

func attrsToString(attrs []html.Attribute) string {
	var fmtdAttrs []string
	for _, attr := range attrs {
		attrStr := fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val)
		fmtdAttrs = append(fmtdAttrs, attrStr)
	}
	return strings.Join(fmtdAttrs, " ")
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild != nil {
			depth--
			fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

// forEachNode calls the functions pre(x) and post(x) for each node x in the tree rooted at n. Both
// functions are optional. pre is called before the children are visited (preorder) and post is
// called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
