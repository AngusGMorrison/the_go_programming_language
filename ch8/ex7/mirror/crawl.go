package mirror

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// crawl saves the page found at link and returns all links present on the page.
func (m *Mirror) crawl(link *url.URL) []*url.URL {
	fmt.Printf("Crawling %s...\n", link)

	if !m.canDownload(link) {
		return nil // skip page
	}

	page, err := getAndParse(link)
	if err != nil {
		log.Println(err)
	}

	foundLinks := m.processLinks(page, link.Path)
	m.save(link, page)
	return foundLinks
}

// canDownload returns false if the link is inaccessible or the page size exceeds m.maxDownload.
func (m *Mirror) canDownload(link *url.URL) bool {
	size, err := downloadSize(link)
	if err != nil {
		log.Printf("sizing %s: %v\n", link, err)
		return false
	}
	if size > m.maxDownload {
		log.Printf("%s too large to download (%d B)\n", link, size)
		return false
	}
	return true
}

// downloadSize returns the size of the requested resource without downloading it.
func downloadSize(link *url.URL) (uint64, error) {
	resp, err := http.Head(link.String())
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return 0, err
	}

	bytes, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	return uint64(bytes), nil
}

// getAndParse downloads the specified URL and parses it into an *html.Node tree.
func getAndParse(link *url.URL) (*html.Node, error) {
	resp, err := http.Get(link.String())
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %v", link, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetching %s: %s", link, resp.Status)
	}

	page, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %v", link, err)
	}
	return page, nil
}

// processLinks extracts all links in an *html.Node tree and replaces them with links to the
// locally mirrored page before returning the originals.
func (m *Mirror) processLinks(page *html.Node, currentPath string) []*url.URL {
	var links []*url.URL
	extractAndReplace := func(n *html.Node) {
		if !(n.Type == html.ElementNode && n.Data == "a") {
			return
		}

		for i, attr := range n.Attr {
			if attr.Key != "href" {
				continue
			}
			foundLink, err := url.Parse(attr.Val)
			if err != nil {
				continue // ignore bad URLs
			}
			if !foundLink.IsAbs() {
				m.completeURL(foundLink, currentPath)
			}
			if foundLink.Hostname() != m.entrypoint.Host {
				continue // ignore URLs outside host
			}
			links = append(links, foundLink)
			// Update links within host domain to refer to local copies.
			n.Attr[i].Val = m.buildLocalPath(foundLink)
		}
	}
	forEachNode(page, extractAndReplace)
	return links
}

// completeURL takes relative URLs and makes them absolute based on the Mirror's host
func (m *Mirror) completeURL(link *url.URL, currentPath string) {
	link.Scheme = m.entrypoint.Scheme
	link.Host = m.entrypoint.Host
	if !strings.HasPrefix(link.Path, "/") {
		// If URL not relative to server root, include path of current page.
		link.Path = currentPath + link.Path
	}
}

// forEachNode iterates over an *html.Node tree and applies a function to each node.
func forEachNode(n *html.Node, f func(n *html.Node)) {
	f(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, f)
	}
}
