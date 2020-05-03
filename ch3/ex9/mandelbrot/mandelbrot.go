package mandelbrot

import (
	"image"
	"image/color"
	"math/cmplx"
)

// Draw creates an image of the Mandelbrot set generated from its arguments
func Draw(settings map[string]float64) *image.RGBA {
	var width, height, zoom = settings["width"], settings["height"], settings["zoom"]
	var xmin, ymin, xmax, ymax = -zoom, -zoom, +zoom, +zoom
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := 0; py < int(height); py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < int(width); px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z
			img.Set(px, py, calcColor(z))
		}
	}
	return img
}

func calcColor(z complex128) color.Color {
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
