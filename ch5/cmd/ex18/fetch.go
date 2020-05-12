// Without changing its behavior, write the fetch function to use defer to close the writable file.
package fetch

import (
	"io"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns the name and length of the local file.
func Fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		if err == nil {
			err = f.Close()
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
