// Many GUIs provide a table widget with a stateful multi-tier sort: the primary sort key is the
// most recently clicked column head, the secondary sort key is the second-most recently clicked
// column head, and so on. Define an implementation of sort.Interface for use by such a table.
// Compare that approach with repeated sorting using sort.Stable.
package main

import (
	"fmt"
	"the_go_programming_language/ch7/cmd/ex8/track"
	"time"
)

var tracks = []*track.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Made Up", 2018, length("2m21s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func main() {
	// No sort order specified
	track.SortBy(tracks, []string{})
	printTracks()

	// Sort with title as the primary key
	sortOrder := []string{"Title", "Artist", "Year"}
	track.SortBy(tracks, sortOrder)
	printTracks()

	sortOrder = []string{"Artist", "Length"}
	track.SortBy(tracks, sortOrder)
	printTracks()

	sortOrder = []string{"Artist", "Year", "Length"}
	track.SortBy(tracks, sortOrder)
	printTracks()
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks() {
	for _, track := range tracks {
		fmt.Printf("%+v\n", track)
	}
	fmt.Println()
}
