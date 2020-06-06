// Package mirror supports the creation of a local copy of a website by crawling the host specfied
// in a given starting URL and saving pages to disk.
package mirror

import (
	"net/url"
)

// Mirror holds the parameters and methods required to crawl a host and download its pages.
type Mirror struct {
	startURL    *url.URL
	host        string
	maxDownload uint64
	rootDir     string
}

// New returns a pointer to a configured Mirror ready to crawl and save the site specified by its
// startURL.
func New(startURL *url.URL, maxDownload uint64, rootDir string) *Mirror {
	return &Mirror{startURL, startURL.Hostname(), maxDownload, rootDir}
}
