package track

import (
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

func (x TrackMultiSort) Len() int           { return len(x.t) }
func (x TrackMultiSort) Less(i, j int) bool { return x.lessOrEqual(x.t[i], x.t[j], x.cols) }
func (x TrackMultiSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// SortBy sorts tracks by the fields named in cols, in order of precedence. I.e. cols[0] is the
// primary sort field, cols[1] the secondary, and so on.
func SortBy(tracks []*Track, cols []string) {
	sort.Sort(multiSort{tracks, cols, func(x, y *Track, cols []string) bool {
		for _, col := range cols {
			switch col {
			case "Title":
				// <= ensures elem won't be swapped if has the same value as the elem it's
				// compared to. The comparison is then made between secondary sort fields.
				// Equivalent to if x.Title != y.Title.
				return x.Title <= y.Title
			case "Artist":
				return x.Artist <= y.Artist
			case "Album":
				return x.Album <= y.Album
			case "Year":
				return x.Year <= y.Year
			case "Length":
				return x.Length <= y.Length
			}
		}
		return false
	}})
}
