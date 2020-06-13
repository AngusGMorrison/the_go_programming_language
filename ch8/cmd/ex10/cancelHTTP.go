// HTTP requests may be cancelled by closing the optional Cancel chanel in the http.Request struct.
// Modify the web crawler of Section 8.6 to support cancellation.
//
// Hint: the http.Get convenience function does not give you an opportunity to customize a Request.
// Instead, create the request using http.NewRequest, set its Cancel field, the perform the request
// by calling http.DefaultClient.Do(req).
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

var done = make(chan struct{})

func main() {
	worklist := make(chan []string) // lists of links returned by crawlers; may contains duplicates
	readyList := make(chan string)  // deduped list of links ready to crawl
	var n int                       // number of pending sends to worklist

	// Listen for cancellation
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// Start with the command-line args.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Start fixed number of goroutines.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range readyList {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				readyList <- link
			}
		}
	}
}

func crawl(link string) []string {
	fmt.Println(link)

	// Skip invalid URLs.
	parsedURL, err := url.Parse(link)
	if err != nil {
		log.Printf("parsing %s: %v\n", link, err)
		return nil
	}

	// Fetch page and return links.
	doc, err := getHTML(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	links := extractLinks(doc, parsedURL)
	return links
}

func getHTML(url string) (*html.Node, error) {
	// Configure cancellable request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GET req for %s: %v", url, err)
	}
	req.Cancel = done

	// Fetch remote data
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getting %s: %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// Parse resp
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return doc, nil
}

func extractLinks(doc *html.Node, address *url.URL) []string {
	var links []string
	extract := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := address.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, extract)
	return links
}

func forEachNode(n *html.Node, f func(*html.Node)) {
	if f != nil {
		f(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, f)
	}
}
