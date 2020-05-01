/*
Modify the Lissajous program to produce images in multiple colors by adding
more values to palette and then displaying them by changing the third argument
of SetColorIndex in some interesting way.
*/

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var cyan = color.RGBA{0x00, 0xFF, 0xFF, 0xFF}
var magenta = color.RGBA{0xFF, 0x00, 0xFF, 0xFF}
var yellow = color.RGBA{0xFF, 0xFF, 0x00, 0xFF}
var palette = []color.Color{color.Black, cyan, magenta, yellow}

const (
	blackIndex   = 0 // First color in palette
	cyanIndex    = 1 // Second color in palette
	magentaIndex = 2
	yellowIndex  = 3
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0 // Relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(i/3) + 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
