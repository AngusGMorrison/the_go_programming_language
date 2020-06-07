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
	list, err := m.saveAndExtractLinks(link)
	if err != nil {
		log.Println(err)
	}
	return list
}

// saveAndExtractLinks downloads valid pages to locations mirroring the host's directory structure,
// and returns a list of URLs found on the page to the crawler.
func (m *Mirror) saveAndExtractLinks(link *url.URL) ([]*url.URL, error) {
	// If page is too large to download, skip it
	size, err := downloadSize(link)
	if err != nil {
		log.Printf("saving %s: %v\n", link, err)
	}
	if size > m.maxDownload {
		log.Printf("%s too large to download (%d B)\n", link, size)
		return nil, nil
	}

	resp, err := http.Get(link.String())
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %v", link, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetching %s: %s", link, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %v", link, err)
	}
	resp.Body.Close() // Close body before writing new file to free fd

	foundLinks, err := m.extractAndReplaceLinks(doc, link.Path)
	if err != nil {
		return foundLinks, fmt.Errorf("extracting links from %q: %v", link, err)
	}
	m.save(link, doc)
	return foundLinks, nil
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

// extractAndReplace links extracts all links in a tree of html Nodes and replaces them with links
// to the equivalent mirrored page locally.
func (m *Mirror) extractAndReplaceLinks(doc *html.Node, currentPath string) ([]*url.URL, error) {
	var links []*url.URL
	extractAndReplace := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}

				link, err := url.Parse(attr.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if !link.IsAbs() {
					link.Scheme = m.startURL.Scheme
					link.Host = m.host
					if !strings.HasPrefix(link.Path, "/") {
						link.Path = currentPath + link.Path
					}
				}

				if link.Hostname() != m.host {
					continue // ignore URLs outside host
				}
				links = append(links, link)
				// Update links within m.host to refer to local copies
				n.Attr[i].Val = m.buildLocalPath(link)
			}
		}
	}
	forEachNode(doc, extractAndReplace)
	return links, nil
}

func forEachNode(n *html.Node, f func(n *html.Node)) {
	f(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, f)
	}
}
