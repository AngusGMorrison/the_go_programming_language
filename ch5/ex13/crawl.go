// Modify crawl to make local copies of the pages it finds, creating directories as necessary.
// Don't make copies of pages that come from a different domain. For example, if the original page
// comes from golang.org, save all files from there, but exclude ones from vimeo.com.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"the_go_programming_language/ch5/links"
)

const savePath = "./crawled_pages"

var targetHosts = make(map[string]bool)
var savedPages = make(map[string]bool)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url [url2 url3...]\n", os.Args[0])
		os.Exit(1)
	}
	// Preserve the target hosts of each input URL to match links that should be downloaded
	if err := captureTargetHosts(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", os.Args[0], err)
		os.Exit(1)
	}
	// Create a directory to hold downloaded pages. Report any error that isn't related to the dir
	// already existing.
	if err := os.Mkdir(savePath, 0777); err != nil && !os.IsExist(err) {
		fmt.Fprintf(os.Stderr, "%s: couldn't create %s dir: %v\n", os.Args[0], savePath, err)
		os.Exit(1)
	}
	// Crawl the web breadth-first, starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

func captureTargetHosts(inputURLs []string) error {
	for _, rawURL := range inputURLs {
		u, err := url.Parse(rawURL)
		if err != nil {
			return fmt.Errorf("parsing %s: %v", rawURL, err)
		}
		targetHosts[u.Host] = true
	}
	return nil
}

// breadthFirst calls f for each item in the worklist. Any items return by f are added to the
// worklist. f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	err := createLocalCopy(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "crawling %s: %v", url, err)
	}
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// createLocalCopy saves a copy of each crawled page that belongs to one of the target hosts and
// has not previously been saved. The directory structure created mirrors the URL path.
func createLocalCopy(rawURL string) error {
	// Determine whether URL should be saved by checking whether it has a target host
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: parsing %s: %v", os.Args[0], rawURL, err)
		return nil
	}
	if !targetHosts[u.Host] {
		return nil
	}

	// Generate the path of the page to be saved, check it hasn't already been saved, then create
	// the directory structure to hold the page.
	localPath := generateLocalPath(u)
	if savedPages[localPath] {
		return nil
	}

	if err = os.MkdirAll(filepath.Dir(localPath), 0777); err != nil {
		log.Fatal(err)
	}

	// Fetch the page contents and save it to disk
	page, err := fetch(rawURL)
	if err != nil {
		return err
	}
	savePage(localPath, page)
	return nil
}

// generateLocalPath creates the path to save the crawled page to. If a filename exists, it is
// preserved, otherwise the page is saved as "index.html".
func generateLocalPath(u *url.URL) string {
	var localPath string
	hasExtension, _ := regexp.MatchString(`/[^/]+\.[^/]+$`, u.Path)
	if hasExtension {
		localPath = u.Path
	} else {
		localPath = u.Path + "/index.html"
	}
	return savePath + "/" + u.Host + localPath
}

// fetch makes an HTTP GET request to the specified URL and returns the page contents
func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s responded %s", url, resp.Status)
	}

	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading from %s: %v", url, err)
	}
	return page, nil
}

func savePage(localPath string, page []byte) {
	file, err := os.Create(localPath)
	defer file.Close()
	if err != nil {
		log.Fatalf("creating file %s: %v", localPath, err)
	}

	_, err = file.Write(page)
	if err != nil {
		log.Fatalf("writing file %s: %v", localPath, err)
	}
}
