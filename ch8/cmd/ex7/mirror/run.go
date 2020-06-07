package mirror

import (
	"fmt"
	"net/url"
	"os"
)

// Run crawls and downloads Mirror's host, mimicking its directory structure locally.
func (m *Mirror) Run() error {
	// Create root directory to save crawled pages to
	if err := os.Mkdir(m.rootDir, 0777); err != nil {
		return fmt.Errorf("couldn't create dir %s: %v", m.rootDir, err)
	}

	var pending int                    // number of links waiting to be crawled
	worklist := make(chan []*url.URL)  // lists of URLs received from crawlers; may have duplicates
	unseenLinks := make(chan *url.URL) // de-duped URLs with the specified host

	// Add starting URL to unseenLinks.
	pending++
	go func() { unseenLinks <- m.startURL }()

	// Create 20 crawler routines to download each unseen page and extract links.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := m.crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// Dedup worklist items, remove links to unwanted hosts, and send remaining unseen links to
	// crawlers.
	seen := make(map[*url.URL]bool)
	for ; pending > 0; pending-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				pending++
				unseenLinks <- link
			}
		}
	}

	return nil
}
