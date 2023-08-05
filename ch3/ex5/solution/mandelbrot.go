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
)

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
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
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
			c := palette[n % uint8(len(palette))]

			r, g, b, a := c.RGBA()
			r -= contrast * uint32(n)
			g -= contrast * uint32(n)
			b -= contrast * uint32(n)
			a -= contrast * uint32(n)

			return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		}
	}
	return color.Black
}

//!-
