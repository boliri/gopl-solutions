package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"

	"surface"
)

func main() {
	var fn string
	if len(os.Args) > 1 {
		fn = os.Args[1]
	}

	var f surface.PlotFn
	switch fn {
		case "eggbox":
			f = surface.Eggbox
		case "saddle":
			f = surface.Saddle
		case "mogguls":
			f = surface.Mogguls
		default:
			f = surface.F
	}

	fmt.Println("Generating SVG representing a 3D surface with the following parameters:")
	fmt.Printf("\tCanvas size: %dx%d\n", surface.Width, surface.Height)
	fmt.Printf("\tCells in grid: %d\n", surface.Cells)
	fmt.Printf("\tRange for XY axes: (-%[1]f, +%[1]f)\n", surface.XYrange)
	fmt.Printf("\tPixels per X or Y unit: %f\n", surface.XYscale)
	fmt.Printf("\tPixels per Z unit: %f\n", surface.Zscale)
	fmt.Printf("\tAngle of X and Y axes: %f\n", surface.Angle)
	fmt.Printf("\tPlot function: %s\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	fmt.Println()

	fmt.Println("Output:")
	surface.GenerateSVG(f)
}