// Use the html/template package (§4.6) to replace printTracks with a function that displays the
// tracks as an HTML table. Use the solution to the previous exercise to arrange that each click on
// a column head makes an HTTP request to sort the table.
package main

import (
	"html/template"
	"log"
	"net/http"
	"the_go_programming_language/ch7/ex8/track"
)

var tableData = []*track.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, track.Length("3m38s")},
	{"Go", "Moby", "Made Up", 2018, track.Length("2m21s")},
	{"Go", "Moby", "Moby", 1992, track.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, track.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, track.Length("4m24s")},
}

const templ = `<html lang="en">
<head></head>
<body>
	<table>
		<thead>
			<tr>
				<th><a href="/?primary_key=Title">Title</a></th>
				<th><a href="/?primary_key=Artist">Artist</a></th>
				<th><a href="/?primary_key=Album">Album</a></th>
				<th><a href="/?primary_key=Year">Year</a></th>
				<th><a href="/?primary_key=Length">Length</a></th>
			</tr>
		</thead>
		<tbody>
		 {{range .}}
			 <tr>
				 <td>{{.Title}}</td>
				 <td>{{.Artist}}</td>
				 <td>{{.Album}}</td>
				 <td>{{.Year}}</td>
				 <td>{{.Length}}</td>
			 </tr>
		 {{end}}
		</tbody>
	</table>
</body>
</html>`

var table = template.Must(template.New("tableHTML").Parse(templ))

// sortOn acts as a rudimentary queue, with the most recently clicked column being added to the
// start of the list.
var sortOn = make([]string, 2)

func main() {
	http.HandleFunc("/", handleSort)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleSort(w http.ResponseWriter, r *http.Request) {
	sortKey, ok := r.URL.Query()["primary_key"]
	if ok {
		sortOn[1], sortOn[0] = sortOn[0], sortKey[0]
	}
	track.SortBy(tableData, sortOn)
	if err := table.Execute(w, tableData); err != nil {
		log.Fatal(err)
	}
}
