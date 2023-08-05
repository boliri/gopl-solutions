// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"fmt"
	"math"
)

const (
	Width, Height = 600, 320            // canvas size in pixels
	Cells         = 100                 // number of grid cells
	XYrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	XYscale       = Width / 2 / XYrange // pixels per x or y unit
	Zscale        = Height * 0.4        // pixels per z unit
	Angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(Angle), math.Cos(Angle) // sin(30°), cos(30°)

func GenerateSVG() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", Width, Height)
	for i := 0; i < Cells; i++ {
		for j := 0; j < Cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if anyInf(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := XYrange * (float64(i)/Cells - 0.5)
	y := XYrange * (float64(j)/Cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := Width/2 + (x-y)*cos30*XYscale
	sy := Height/2 + (x+y)*sin30*XYscale - z*Zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func anyInf(nums ...float64) bool {
    for _, n := range nums {
		if math.IsInf(n, 0) {
			return true
		}
	}
	return false
}

//!-
