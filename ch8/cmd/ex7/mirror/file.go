package mirror

import (
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"

	"golang.org/x/net/html"
)

func (m *Mirror) save(link *url.URL, doc *html.Node) {
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

	err = html.Render(file, doc)
	if err != nil {
		log.Fatalf("copying to %q: %v", path, err)
	}
}

// buildLocalPath constructs a local filepath from the current URL
func (m *Mirror) buildLocalPath(link *url.URL) string {
	var filename string
	if hasExt, _ := regexp.MatchString(`/[^/]+\.[a-zA-Z]+$`, link.Path); hasExt {
		filename = link.Path
	} else {
		filename = link.Path + "/index.html"
	}
	return m.rootDir + "/" + filename
}
