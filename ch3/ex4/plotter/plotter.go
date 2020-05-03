/*
Math plotting functions, constants and variables.
*/

package plotter

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var xyscale float64
var zscale float64
var width float64
var height float64
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// Plot a mathematical function onto an SVG canvas
func Plot(settings map[string]string) (string, error) {
	var err error
	width, height, err = parseDimensions(settings["width"], settings["height"])
	if err != nil {
		return "", err
	}
	xyscale = width / 2 / xyrange
	zscale = height * 0.4

	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: white; fill: #%s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", settings["color"], settings["width"], settings["height"])
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			svg += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	svg += fmt.Sprintf("</svg>")
	return svg, nil
}

func parseDimensions(widthStr string, heightStr string) (float64, float64, error) {
	width, widthErr := strconv.ParseFloat(widthStr, 64)
	height, heightErr := strconv.ParseFloat(heightStr, 64)
	if widthErr != nil || heightErr != nil || width < 0 || height < 0 {
		errMsg := fmt.Sprintf("Invalid dimensions: width %s, height %s", widthStr, heightStr)
		return 0, 0, errors.New(errMsg)
	}
	return width, height, nil
}

func corner(i, j int) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Distance from (0,0)
	return math.Sin(r) / r
}
