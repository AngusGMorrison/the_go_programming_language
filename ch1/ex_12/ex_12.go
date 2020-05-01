/*
Modify the Lissajous server to read parameter values from the URL. For example,
you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets
the number of cycles to 20 instead of the default 5. Use the default
strconv.Atoi function to convert the string parameter into an integer. You can
see its documentation with go doc strconv.Atoi.
*/

package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	settings, err := parseUserSettings(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	lissajous(w, settings)
}

func parseUserSettings(r *http.Request) (map[string]int, error) {
	var settings = map[string]int{
		"cycles":  5,
		"size":    100,
		"nframes": 64,
		"delay":   8,
	}

	// Map over query parameters, adding them to the settings map and overriding
	// existing values.
	for k, v := range r.URL.Query() {
		intVal, err := strconv.Atoi(v[0])
		if err != nil {
			errMsg := fmt.Sprintf("Invalid parameter %s: %s", k, v)
			return nil, errors.New(errMsg)
		}
		settings[k] = intVal
	}
	return settings, nil
}

func lissajous(w http.ResponseWriter, settings map[string]int) {
	palette := []color.Color{color.White, color.Black}

	const (
		whiteIndex = 0
		blackIndex = 1
		res        = 0.001
	)

	freq := rand.Float64() * 3.0 // Relative frequency of y oscillator
	anim := gif.GIF{LoopCount: settings["nframes"]}
	phase := 0.0 // Phase difference

	for i := 0; i < settings["nframes"]; i++ {
		size := settings["size"]
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		cycles := float64(settings["cycles"])
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, settings["delay"])
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
