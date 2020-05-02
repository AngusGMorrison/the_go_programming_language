/*
If the function f returns a non-finite float64 value, the SVG file will contain
invalid <polygon> elements (although many SVG renderers handle this gracefully).
Modify the program to skip invalid polygons.
*/

package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 //number of grid cells
	xyrange       = 30.0                //axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angus of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	out, err := os.Create("plot.svg")
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, acol := corner(i+1, j)
			bx, by, bcol := corner(i, j)
			cx, cy, ccol := corner(i, j+1)
			dx, dy, dcol := corner(i+1, j+1)
			red, blue := calcPolygonColor(acol, bcol, ccol, dcol)
			fmt.Fprintf(out, "<polygon fill='rgba(%d, 0, %d, 1)' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				red, blue, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
	out.Close()
}

type colorInfo struct {
	blue, red int
}

func corner(i, j int) (float64, float64, colorInfo) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z and calculate color
	z := f(x, y)
	color := calcCornerColor(z)

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, color
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Distance from (0,0)
	return math.Sin(r) / r
}

func calcCornerColor(z float64) colorInfo {
	var col colorInfo
	if z > 0 {
		col.red = 255
	} else {
		col.blue = 255
	}
	return col
}

func calcPolygonColor(a, b, c, d colorInfo) (int, int) {
	red := (a.red + b.red + c.red + d.red) / 4
	blue := (a.blue + b.blue + c.blue + d.blue) / 4
	return red, blue
}
