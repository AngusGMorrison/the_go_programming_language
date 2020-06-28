package mirror

import (
	"fmt"
	"net/url"
	"os"
)

// Run crawls and downloads the Mirror's entrypoint host, mimicking its directory structure locally.
func (m *Mirror) Run() error {
	// Create root directory to save crawled pages to. If this operation fails (e.g. the program
	// doesn't have permissions for the target dir), we want to know about it before crawling.
	if err := os.Mkdir(m.rootDir, 0777); err != nil {
		return fmt.Errorf("couldn't create dir %s: %v", m.rootDir, err)
	}

	var pending int                    // number of links waiting to be crawled. Run returns at 0
	worklist := make(chan []*url.URL)  // lists of URLs received from crawlers; may have duplicates
	unseenLinks := make(chan *url.URL) // de-duped URLs

	// Add starting URL to unseenLinks via goroutine to avoid deadlock.
	pending++
	go func() { unseenLinks <- m.entrypoint }()

	// Create 20 crawler routines to download each unseen page and extract links.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := m.crawl(link)
				go func() { worklist <- foundLinks }() // send via goroutine to avoid deadlock
			}
		}()
	}

	// Dedup worklist links and send to crawlers.
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
