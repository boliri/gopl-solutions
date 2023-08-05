package main

import (
	"fmt"
	"os"

	"mandelbrot"
)

func main() {
	fmt.Println("Generating PNG of a Newton's fractal using parameters:")
	fmt.Printf("\tImage size: %dx%d\n", mandelbrot.Width, mandelbrot.Height)
	fmt.Printf("\tInterval of real numbers: [%d, %d]\n", mandelbrot.Xmin, mandelbrot.Xmax)
	fmt.Printf("\tInterval of imaginary numbers: [%d, %d]\n", mandelbrot.Ymin, mandelbrot.Ymax)
	fmt.Printf("\tMax iterations per point: %d\n", mandelbrot.NewtonIterations)
	fmt.Println()

	fmt.Println("Creating file...")
	f, err := os.Create("../solution/fractal.png")
	defer f.Close()
	if err != nil {
		fmt.Print("Error while creating file: ", err)
		return
	}

	mandelbrot.GeneratePng(f)
	fmt.Println("PNG file was generated")
}