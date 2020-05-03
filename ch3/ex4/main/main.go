/*
Following the approach of the Lissajous example in Section 1.7, construct a web server that
compputes surfaces and writes AVG data to the clinet. The server must set the Content-Type header
like this:

	w.Header().Set("Content-Type", "image/svg+xml")

(This step was not required in the Lissajous example because the server uses standard heuristics to
recognize common formats like PNG from the first 512 bytes of the response, and generates the
proper header.) Allow the client to specify values like height, width, and color as HTTP request
parameters.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"the_go_programming_language/ch3/ex4/plotter"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	settings := getClientSettings(r.URL.Query())
	plot, err := plotter.Plot(settings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, plot)
}

func getClientSettings(params url.Values) map[string]string {
	var settings = map[string]string{
		"width":  "600",
		"height": "320",
		"color":  "grey",
	}

	for k, v := range params {
		settings[k] = v[0]
	}
	return settings
}
