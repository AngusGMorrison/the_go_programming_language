/*
Color each polygon based on its height, so that the peaks are colored red (#ff0000) and the valleys
blue (#0000ff).
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
			finite := true
			ax, ay := corner(i+1, j, &finite)
			bx, by := corner(i, j, &finite)
			cx, cy := corner(i, j+1, &finite)
			dx, dy := corner(i+1, j+1, &finite)
			if finite {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintf(out, "</svg>")
	out.Close()
}

func corner(i, j int, finitePtr *bool) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	// Set the finite flag to false, preventing printing of the <polygon>
	if math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sy) || math.IsNaN(sx) {
		*finitePtr = false
	}
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Distance from (0,0)
	return math.Sin(r) / r
}
