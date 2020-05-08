// The popular web comic xkcd has a JSON interface. For example, a request to
// https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of many
// favorites. Download each URL (once!) amd build an offline index. Write a tool xkcd that,
// using this index, prints the URL and transcript of each comic that matches a search term provided
// on the command line.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"the_go_programming_language/ch4/ex12/xkcd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s search_term\n", os.Args[0])
	}

	if xkcd.ArchiveNotExist() {
		if err := xkcd.CreateArchive(); err != nil {
			log.Fatal(err)
		}
	}

	searchTerm := strings.ToLower(os.Args[1])
	comics, err := xkcd.SearchArchive(searchTerm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matching comics:\n\n")
	for _, comic := range comics {
		fmt.Printf("%s\n%s\n\n", comic.Title, comic.Transcript)
	}
}
