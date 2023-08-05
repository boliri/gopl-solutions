// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"fmt"
	"math"

	"github.com/atotto/clipboard"
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

type PlotFn func(x, y float64) float64

type Polygon struct {
	ax, ay, az float64
	bx, by, bz float64
	cx, cy, cz float64
	dx, dy, dz float64
}

func (p Polygon) toHtml(color string) string {
	return fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style=\"fill:#%s\"/>",
		p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy, color)
}

func (p Polygon) getPeakZ() float64 {
	return max(p.az, p.bz, p.cz, p.dz)
}

func (p Polygon) getValleyZ() float64 {
	return min(p.az, p.bz, p.cz, p.dz)
}

func GenerateSVG(f PlotFn) {
	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: black; stroke-width: 0.7' "+
		"width='%d' height='%d'>", Width, Height)

	var highestPeak, lowestValley float64
	var polygons []Polygon

	for i := 0; i < Cells; i++ {
		for j := 0; j < Cells; j++ {
			ax, ay, az := corner(i+1, j, f)
			bx, by, bz := corner(i, j, f)
			cx, cy, cz := corner(i, j+1, f)
			dx, dy, dz := corner(i+1, j+1, f)

			if anyInf(ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz) {
				continue
			}

			polygon := Polygon{ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz}
			polygons = append(polygons, polygon)

			polygonPeak := polygon.getPeakZ()
			polygonValley := polygon.getValleyZ()
			if polygonPeak > highestPeak {
				highestPeak = polygonPeak
			} else if polygonPeak < lowestValley {
				lowestValley = polygonValley
			}
		}
	}

	for _, p := range polygons {
		polygonPeak := p.getPeakZ()
		polygonValley := p.getValleyZ()

		var colorHex string
		if math.Abs(polygonPeak) > math.Abs(polygonValley) {
			red := int(math.Abs((polygonPeak / highestPeak) * 255))
			colorHex = fmt.Sprintf("%2x0000", red)
		} else {
			blue :=  int(math.Abs((polygonValley / lowestValley) * 255))
			colorHex = fmt.Sprintf("0000%2x", blue)
		}

		svg += fmt.Sprintln(p.toHtml(colorHex))
	}

	svg += fmt.Sprintln("</svg>")
	fmt.Print(svg)

	// Write to clipboard to ease testing output
	clipboard.WriteAll(svg)
}

func corner(i, j int, f PlotFn) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := XYrange * (float64(i)/Cells - 0.5)
	y := XYrange * (float64(j)/Cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := Width/2 + (x-y)*cos30*XYscale
	sy := Height/2 + (x+y)*sin30*XYscale - z*Zscale
	return sx, sy, z
}

func F(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func Eggbox(x, y float64) float64 {
	return -0.1 * (math.Sin(x) + math.Sin(y))
}

func Saddle(x, y float64) float64 {
	return math.Pow(x, 2) / math.Pow(25, 2) - math.Pow(y, 2) / math.Pow(17, 2)
}

func Mogguls(x, y float64) float64 {
	return -0.1 * (math.Sin(x / 2) + math.Sin(y / 2))
}

func anyInf(nums ...float64) bool {
    for _, n := range nums {
		if math.IsInf(n, 0) {
			return true
		}
	}
	return false
}

func max(nums ...float64) float64 {
	var res float64
	for i := 0; i < len(nums) - 1; i++ {
		res = math.Max(nums[i], nums[i+1])
	}
	return res
}

func min(nums ...float64) float64 {
	var res float64
	for i := 0; i < len(nums) - 1; i++ {
		res = math.Min(nums[i], nums[i+1])
	}
	return res
}

//!-
