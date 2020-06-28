// Following the approach of mirroredQuery in Section 8.4.4, implement a variant of fetch that
// requests several URLs concurrently. As soon as the first response arrives, cancel the other
// requests.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	done := make(chan struct{})
	responses := make(chan string, len(os.Args[1:]))

	for _, url := range os.Args[1:] {
		go func(url string) {
			body, err := fetch(url, done)
			if err != nil {
				log.Printf("fetching %s: %v\n", url, err) // log cancellations
				return
			}
			responses <- string(body)
			close(done)
		}(url)
	}

	fmt.Printf("%s\n", <-responses)
}

func fetch(url string, done <-chan struct{}) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GET request: %v", err)
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}
	return body, nil
}
