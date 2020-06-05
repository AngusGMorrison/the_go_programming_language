// Take an existing CPU-bound sequential program, such as the Mandelbrot program of Section 3.3 or
// the 3D surface computation of Section 3.2, and execute its main loop in parallel using channels
// for communication. How much faster does it run on a multiprocessor machine? What is the optimal
// number of goroutines to use?
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	file, err := os.Create("mandelbrot.png")
	if err != nil {
		log.Fatal(err)
	}
	img := buildImage(5)
	if err := png.Encode(file, img); err != nil {
		log.Fatal(err)
	}
}

func buildImage(maxGortns int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	yPixels := make(chan float64)
	var wg sync.WaitGroup
	for i := 0; i < maxGortns; i++ {
		wg.Add(1)
		go buildColumn(yPixels, img, &wg)
	}
	for py := 0.0; py < height; py++ {
		yPixels <- py
	}
	close(yPixels)
	wg.Wait()
	return img
}

func buildColumn(yPixels <-chan float64, img *image.RGBA, wg *sync.WaitGroup) {
	defer wg.Done()
	for py := range yPixels {
		y := py/height*(ymax-ymin) + ymin
		for px := 0.0; px < width; px++ {
			x := px/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px,py) represents complex value z.
			img.Set(int(px), int(py), mandelbrot(z))
		}
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
