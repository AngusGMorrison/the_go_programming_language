package mirror

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// crawl saves the page found at link and returns all links present on the page.
func (m *Mirror) crawl(link *url.URL) []*url.URL {
	fmt.Println("Crawling %s...", link)
	list, err := m.saveAndExtractLinks(link)
	if err != nil {
		log.Println(err)
	}
	return list
}

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

}

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
