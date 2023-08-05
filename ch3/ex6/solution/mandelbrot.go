// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// This version of Mandelbrot emits a PNG image of the Mandelbrot fractal with a sampling technique applied to reduce pixelation.
package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"math/cmplx"
)

const (
	Xmin, Ymin, Xmax, Ymax = -2, -2, +2, +2
	Width, Height          = 1024, 1024
	SubpixelsPerPixel	   = 4                      // number of subpixels to calculate color
	SubpixelsPerAxis	   = SubpixelsPerPixel / 2  // number of pixels to generate per axis (X or Y)
	SubpixelReductionRate  = 10000                  // adjusts the position of a subpixel so it's closer to the original pixel
)

func GeneratePng(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	for py := 0; py < Height; py++ {
		for px := 0; px < Width; px++ {
			var r, g, b, a float64

			for i := 1.0; i <= SubpixelsPerAxis; i++ {
				xfactor := math.Pow(-1, i) * (i / SubpixelsPerAxis / SubpixelReductionRate)
				x := (float64(px) + xfactor) / Width * (Xmax-Xmin) + Xmin

				for j := 1.0; j <= SubpixelsPerAxis; j++ {
					yfactor := math.Pow(-1, j) * (j / SubpixelsPerAxis / SubpixelReductionRate)
					y := (float64(py) + yfactor) / Height * (Ymax-Ymin) + Ymin

					// Image point (px, py) represents complex value z.
					subz := complex(x, y)

					subr, subg, subb, suba := mandelbrot(subz).RGBA()
					r += float64(subr) / float64(SubpixelsPerPixel)
					g += float64(subg) / float64(SubpixelsPerPixel)
					b += float64(subb) / float64(SubpixelsPerPixel)
					a += float64(suba) / float64(SubpixelsPerPixel)
				}
			}

			c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
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

//!-
