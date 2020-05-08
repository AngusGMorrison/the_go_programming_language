// Package xkcd allows retrieval, archival and search of xkcd webcomics.
package xkcd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Comic holds basic informat about a single xkcd comic
type Comic struct {
	Num        int
	Title      string
	Transcript string
}

const (
	archive = "./archive.json"
	// Limit maxComics to archive to avoid throttling
	maxComics = 404
	xkcdURL   = "http://xkcd.com/%d/info.0.json"
)

// ArchiveNotExist reports the presence of a local xkcd archive file
func ArchiveNotExist() bool {
	_, err := os.Stat(archive)
	return os.IsNotExist(err)
}

// CreateArchive adds every xkcd JSON representation to a local JSON file
func CreateArchive() error {
	fmt.Printf("Creating archive. This may take some time...\n\n")
	file, err := os.Create(archive)
	if err != nil {
		return fmt.Errorf("Unable to create archive at %s", archive)
	}

	defer file.Close()

	var comics []*Comic
	for i := 1; i <= maxComics; i++ {
		c, err := getComic(i)
		if err != nil {
			switch err.(type) {
			case *url.Error:
				log.Fatal(err)
			default:
				fmt.Fprintf(os.Stderr, "%s\n", err.Error())
				continue
			}
		}
		comics = append(comics, c)
	}

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(comics); err != nil {
		log.Fatal(err)
	}
	return nil
}

func getComic(num int) (*Comic, error) {
	url := fmt.Sprintf(xkcdURL, num)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Comic %d responded with %d", num, resp.StatusCode)
	}

	var c Comic
	if err = json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

// SearchArchive scans the local xkcd archive for comics with titles containing searchTerm.
// Hits are added to a slice of matching comics. Note, this sort of linear search is extremely
// inefficient, but it serves the purpose of the excercise in the absence of a DBMS.
func SearchArchive(searchTerm string) ([]*Comic, error) {
	file, err := os.Open(archive)
	if err != nil {
		return []*Comic{}, fmt.Errorf("Unable to open archive at %s", archive)
	}

	defer file.Close()

	var matchingComics []*Comic
	decoder := json.NewDecoder(file)

	// read open bracket
	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for decoder.More() {
		var c Comic
		// decode an array value
		err = decoder.Decode(&c)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(strings.ToLower(c.Title), searchTerm) {
			matchingComics = append(matchingComics, &c)
		}
	}

	return matchingComics, nil
}
