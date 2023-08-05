// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// This version of Mandelbrot emits a PNG image of the Mandelbrot fractal in a colored fashion.
package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

const (
	Xmin, Ymin, Xmax, Ymax = -2, -2, +2, +2
	Width, Height          = 1024, 1024

	NewtonIterations       = 37             // max number of points tested to find a root for z^4 - 1
)

var roots [4]complex128 = [4]complex128{
	complex(1, 0),
	complex(-1, 0),
	complex(0, 1),
	complex(0, -1),
}

var palette []color.Color = []color.Color{
	color.RGBA{0, 0, 255, 255},   // blue
	color.RGBA{0, 255, 0, 255},   // green
	color.RGBA{255, 0, 255, 255}, // magenta
	color.RGBA{255, 255, 0, 255}, // yellow
}

func GeneratePng(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	for py := 0; py < Height; py++ {
		y := float64(py)/Height*(Ymax-Ymin) + Ymin
		for px := 0; px < Width; px++ {
			x := float64(px)/Width*(Xmax-Xmin) + Xmin

			// Image point (px, py) represents complex value z.
			z := complex(x, y)

			c := newton(z)
			img.Set(px, py, c)
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
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

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
// returns the color from the palette that is associated to the closest root for z^4 - 1, with some contrast applied (the closest, the highest)
// if the point is not close enough to any root, returns a black color
func newton(z complex128) color.Color {
	var c color.Color = color.Black
	const contrast = 7

	for i := int(0); i < NewtonIterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		for ri, root := range roots {
			distance := cmplx.Abs(z - root)
			if distance <= 1e-6 {
				c := palette[ri]

				r, g, b, a := c.RGBA()
				r -= contrast * uint32(i)
				g -= contrast * uint32(i)
				b -= contrast * uint32(i)
				a -= contrast * uint32(i)

				return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			}
		}
	}
	return c
}
//!-
