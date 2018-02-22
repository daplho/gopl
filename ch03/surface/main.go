// Surface computes an SVG rendering of a 3-D surface function
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 1200, 640           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (outX float64, outY float64, ok bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute suface height z.
	z, ok := f(x, y)
	if !ok {
		return
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	outX = width/2 + (x-y)*cos30*xyscale
	outY = height/2 + (x+y)*sin30*xyscale - z*zscale
	ok = true
	return
}

func f(x, y float64) (value float64, ok bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	result := math.Sin(r) / r
	if math.IsInf(result, 0) {
		return 0.0, false
	}
	return result, true
}
