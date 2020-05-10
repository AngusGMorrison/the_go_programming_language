// Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url\n", os.Args[0])
		os.Exit(1)
	}

	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Words: %d; Images: %d\n", words, images)
}

// CountWordsAndImages does an HTTP GET request for the HTML document url and returns the number of
// words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

// countWordsAndImages recurses to the bottom of a node tree, returning the cumulative total of
// each node's words and images on the way back up.
func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	} else if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cWords, cImages := countWordsAndImages(c)
		words += cWords
		images += cImages
	}

	return
}
