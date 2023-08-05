package main

import (
	"fmt"
	"os"

	"mandelbrot"
)

func main() {
	fmt.Println("Generating PNG of a full-colored Mandelbrot fractal using parameters:")
	fmt.Printf("\tImage size: %dx%d\n", mandelbrot.Width, mandelbrot.Height)
	fmt.Printf("\tInterval of real numbers: [%d, %d]\n", mandelbrot.Xmin, mandelbrot.Xmax)
	fmt.Printf("\tInterval of imaginary numbers: [%d, %d]\n", mandelbrot.Ymin, mandelbrot.Ymax)
	fmt.Println()

	fmt.Println("Creating file for fractal based on 64-bits complex numbers...")
	f, err := os.Create("../solution/fractal_complex64.png")
	defer f.Close()
	if err != nil {
		fmt.Print("Error while creating file: ", err)
		return
	}
	mandelbrot.GeneratePng(f, mandelbrot.MandelbrotComplex64)

	fmt.Println("Creating file for fractal based on 128-bits complex numbers...")
	f, err = os.Create("../solution/fractal_complex128.png")
	if err != nil {
		fmt.Print("Error while creating file: ", err)
		return
	}
	mandelbrot.GeneratePng(f, mandelbrot.MandelbrotComplex128)

	fmt.Println("Creating file for fractal based on big float numbers...")
	f, err = os.Create("../solution/fractal_bigfloat.png")
	if err != nil {
		fmt.Print("Error while creating file: ", err)
		return
	}
	mandelbrot.GeneratePng(f, mandelbrot.MandelbrotBigFloat)

	// This one takes too long as big rational numbers are expensive, so we'll skip it
	// fmt.Println("Creating file for fractal based on big rational numbers...")
	// f, err = os.Create("../solution/fractal_bigrational.png")
	// if err != nil {
	// 	fmt.Print("Error while creating file: ", err)
	// 	return
	// }
	// mandelbrot.GeneratePng(f, mandelbrot.MandelbrotBigRational)

	fmt.Println("PNG files were generated")
}