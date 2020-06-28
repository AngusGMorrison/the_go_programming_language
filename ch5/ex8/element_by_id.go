// Modify forEachNode so that the pre and post functions return a boolean result
// indicating whether to continue traversal. Use it to write a function ElementByID with the
// following signature that finds the first HTML element with teh specified id attribute. The
// function should stop the traversal as soon as a match is found.
//
//		func ElementByID(doc *html.Node, id string) *html.Node
//
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: %s url id", os.Args[0])
		os.Exit(1)
	}
	doc, err := fetchNodes(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	el := elementByID(doc, os.Args[2])
	if el == nil {
		fmt.Printf("%s: id %s not found at %s\n", os.Args[0], os.Args[2], os.Args[1])
	} else {
		printNode(el)
	}
}

// fetchNodes GETs the specified URL and returns its parsed HTML node tree
func fetchNodes(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s responded with status code %d", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %v", url, err)
	}
	return doc, nil
}

func elementByID(doc *html.Node, id string) *html.Node {
	matchID := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					return true
				}
			}
		}
		return false
	}

	return forEachNode(doc, matchID, nil)
}

func forEachNode(n *html.Node, pre, post func(*html.Node) bool) *html.Node {
	if pre != nil {
		if pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if match := forEachNode(c, pre, post); match != nil {
			return match
		}
	}

	if post != nil {
		if post(n) {
			return n
		}
	}

	return nil
}

func printNode(n *html.Node) {
	fmt.Printf("<%s ", n.Data)
	for _, attr := range n.Attr {
		fmt.Printf(`%s="%s"`, attr.Key, attr.Val)
	}
	fmt.Println(">")
}
