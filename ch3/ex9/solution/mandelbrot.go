// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

type MandelbrotConfig struct {
	X			    int
	Y  	            int
	Zoom			int
}

func newDefaultMandelbrotConfig() MandelbrotConfig {
	return MandelbrotConfig{
		X: 2,
		Y: 2,
		Zoom: 1,
	}
}

func NewMandelbrotConfig(x, y, zoom int) MandelbrotConfig {
	cfg := newDefaultMandelbrotConfig()
	cfg.X = x
	cfg.Y = y
	cfg.Zoom = zoom

	return cfg
}

const Width, Height = 1024, 1024

func GeneratePNG(out io.Writer, cfg MandelbrotConfig) {
	xmin, ymin, xmax, ymax := -cfg.X, -cfg.Y, +cfg.X, +cfg.Y

	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	for py := 0; py < Height; py++ {
		y := (float64(py)/float64(Height)*float64(ymax-ymin) + float64(ymin)) / float64(cfg.Zoom)
		for px := 0; px < Width; px++ {
			x := (float64(px)/float64(Width)*float64(xmax-xmin) + float64(xmin)) / float64(cfg.Zoom)
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
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
