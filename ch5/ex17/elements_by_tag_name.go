// Write a variadic function ElementsByTagName that, given an HTML node tree and zero or more names,
// returns all the elements that match one of those names. Here are two example calls:
//
//		func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
//
//		images := ElementsByTagName(doc, "img")
//		headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
//
package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("response %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	matchingNodes := elementsByTagName(doc, "img", "h2")
	for _, node := range matchingNodes {
		fmt.Printf("%+v\n", node)
	}
}

func elementsByTagName(doc *html.Node, names ...string) []*html.Node {
	if len(names) == 0 {
		return []*html.Node{}
	}

	// Create a map from names for O(1) lookup
	targetNames := make(map[string]bool)
	for _, name := range names {
		targetNames[name] = true
	}

	matchingNodes := make([]*html.Node, 0)
	// forEachNode iterates over all nodes in the tree and appends nodes with the names stored in
	// targetNames to matchingNodes
	var forEachNode func(*html.Node)
	forEachNode = func(n *html.Node) {
		if n.Type == html.ElementNode && targetNames[n.Data] {
			matchingNodes = append(matchingNodes, n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c)
		}
	}

	forEachNode(doc)
	return matchingNodes
}
