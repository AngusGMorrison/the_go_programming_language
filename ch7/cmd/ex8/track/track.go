package track

import (
	"fmt"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// multiSort implements sort.Interface using a comparison function that takes a slice of field
// names, cols, and sorts on those fields in order of precedence.
type multiSort struct {
	t           []*Track
	cols        []string
	lessOrEqual func(x, y *Track, cols []string) bool
}

func (x multiSort) Len() int           { return len(x.t) }
func (x multiSort) Less(i, j int) bool { return x.lessOrEqual(x.t[i], x.t[j], x.cols) }
func (x multiSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// SortBy sorts tracks by the fields named in cols, in order of precedence. I.e. cols[0] is the
// primary sort field, cols[1] the secondary, and so on.
func SortBy(tracks []*Track, cols []string) {
	sort.Sort(multiSort{tracks, cols, func(x, y *Track, cols []string) bool {
		for _, col := range cols {
			if col == "Title" && x.Title != y.Title {
				return x.Title < y.Title
			} else if col == "Artist" && x.Artist != y.Artist {
				return x.Artist < y.Artist
			} else if col == "Album" && x.Album != y.Album {
				return x.Album < y.Album
			} else if col == "Year" && x.Year != y.Year {
				return x.Year < y.Year
			} else if col == "Length" && x.Length != y.Length {
				return x.Length <= y.Length
			}
		}
		return true
	}})
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func PrintTracks(tracks []*Track) {
	for _, track := range tracks {
		fmt.Printf("%+v\n", track)
	}
	fmt.Println()
}
