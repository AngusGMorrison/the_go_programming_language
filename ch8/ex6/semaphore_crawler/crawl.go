// Add depth-limiting to the concurrent crawler. That is, if the user sets -depth=3, then only URLs
// reachable by at most three links will be fetched.
package main

import (
	"flag"
	"fmt"
	"log"
	"the_go_programming_language/ch5/links"
)

// depthLtdCrawl contains a slice of links found by a crawler in addition to the depth they were
// found at.
type depthLtdCrawl struct {
	depth int
	links []string
}

var maxDepth int

func init() {
	flag.IntVar(&maxDepth, "depth", 3, "max search depth")
	flag.Parse()
}

func main() {
	worklist := make(chan *depthLtdCrawl)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments
	n++
	go func() { worklist <- &depthLtdCrawl{0, flag.Args()} }()

	// Crawl the web concurrently
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		// Check list depth is valid before marking links as seen - they may be accessible by
		// shorter routes and should not be skipped by the crawler.
		if list.depth >= maxDepth {
			continue
		}
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- &depthLtdCrawl{
						depth: list.depth + 1,
						links: crawl(list.depth, link),
					}
				}(link)
			}
		}
	}
}

// tokens is a counting semaphore used to enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(depth int, url string) []string {
	fmt.Printf("%d %s\n", depth, url)
	tokens <- struct{}{} // acquire a token
	// Note: it is highly recommended to add a timeout to GET requests, or the program may appear
	// to hang when it encounters large downloads
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}
