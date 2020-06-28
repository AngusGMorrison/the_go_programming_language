package mirror

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// save renders a tree of *html.Node to disk.
func (m *Mirror) save(link *url.URL, page *html.Node) {
	path := m.buildLocalPath(link)
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0777); err != nil {
		log.Fatalf("creating dir %q: %v", dir, err)
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("creating file %q: %v", path, err)
	}
	defer file.Close()

	err = html.Render(file, page)
	if err != nil {
		log.Fatalf("copying to %q: %v", path, err)
	}
}

// buildLocalPath constructs a local filepath from the current URL.
func (m *Mirror) buildLocalPath(link *url.URL) string {
	var filename string
	if hasExt, _ := regexp.MatchString(`/[^/]+\.[a-zA-Z]+$`, link.Path); hasExt {
		filename = link.Path
	} else {
		var sep string
		if !strings.HasSuffix(link.Path, "/") {
			sep = "/"
		}
		filename = fmt.Sprintf("%s%sindex.html", link.Path, sep)
	}
	return m.rootDir + filename
}
