/*
Write a web server that renders fractals and writes the image data to the client. Allow the client
to specify the x, y and zoom values as parameters to the HTTP request.
*/

package main

import (
	"image/png"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"the_go_programming_language/ch3/ex9/mandelbrot"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	settings, err := parseClientSettings(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	img := mandelbrot.Draw(settings)
	png.Encode(w, img)
}

func parseClientSettings(params url.Values) (map[string]float64, error) {
	var settings = map[string]float64{
		"width":  1024,
		"height": 1024,
		"zoom":   2,
	}
	for k := range settings {
		if v, found := params[k]; found == true {
			floatVal, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				return nil, err
			}
			settings[k] = floatVal
		}
	}
	return settings, nil
}
