// Many GUIs provide a table widget with a stateful multi-tier sort: the primary sort key is the
// most recently clicked column head, the secondary sort key is the second-most recently clicked
// column head, and so on. Define an implementation of sort.Interface for use by such a table.
// Compare that approach with repeated sorting using sort.Stable.
package main

import (
	"the_go_programming_language/ch7/cmd/ex8/track"
)

var tracks = []*track.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, track.Length("3m38s")},
	{"Go", "Moby", "Made Up", 2018, track.Length("2m21s")},
	{"Go", "Moby", "Moby", 1992, track.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, track.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, track.Length("4m24s")},
}

func main() {
	// No sort order specified
	track.SortBy(tracks, []string{})
	track.PrintTracks(tracks)

	// Sort with title as the primary key
	sortOrder := []string{"Title", "Artist", "Year"}
	track.SortBy(tracks, sortOrder)
	track.PrintTracks(tracks)

	sortOrder = []string{"Artist", "Length"}
	track.SortBy(tracks, sortOrder)
	track.PrintTracks(tracks)

	sortOrder = []string{"Artist", "Year", "Length"}
	track.SortBy(tracks, sortOrder)
	track.PrintTracks(tracks)

	sortOrder = []string{"Artist", "Album"}
	track.SortBy(tracks, sortOrder)
	track.PrintTracks(tracks)
}
