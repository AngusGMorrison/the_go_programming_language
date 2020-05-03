/*
Implement a full-color Mandelbrot set using the function image.NewRGBA and the type color.RGBA or
color.YCbCr.
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}
	out, err := os.Create("mandelbrot.png")
	if err != nil {
		log.Fatal("Unable to create output file")
	}
	png.Encode(out, img) // NOTE: ignoring errors
	out.Close()
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{n, 255 - contrast*n, 123, 255}
		}
	}
	return color.Black
}
