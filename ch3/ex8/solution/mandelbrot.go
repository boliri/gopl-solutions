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
	"math/big"
	"math/cmplx"
)

const (
	Xmin, Ymin, Xmax, Ymax = -2, -2, +2, +2
	Width, Height          = 1024, 1024
)

type mandelbrotFn func(complex128) color.Color

func GeneratePng(out io.Writer, fn mandelbrotFn) {
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	for py := 0; py < Height; py++ {
		y := float64(py) / Height * (Ymax-Ymin) + Ymin
		for px := 0; px < Width; px++ {
			x := float64(px) / Width * (Xmax-Xmin) + Xmin
			z := complex128(complex(x, y))
			// Image point (px, py) represents complex value z.
			img.Set(px, py, fn(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func MandelbrotComplex64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func MandelbrotComplex128(z complex128) color.Color {
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

func MandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	// original z
	zReal := (&big.Float{}).SetFloat64(float64(real(z)))
	zImag := (&big.Float{}).SetFloat64(float64(imag(z)))

	// Nth z (previous iteration)
	prevZreal, prevZimag := &big.Float{}, &big.Float{}

	for n := uint8(0); n < iterations; n++ {
		// z * z = (a + bi) ^ 2 = a^2 + b^2*i^2 + 2abi = a^2 - b^2 + 2abi
		// real = a^2 - b^2
		// imag = 2abi
		currZreal := &big.Float{}
		currZimag := &big.Float{}

		currZreal.Mul(prevZreal, prevZreal).Sub(currZreal, (&big.Float{}).Mul(prevZimag, prevZimag)).Add(currZreal, zReal)
		currZimag.Mul(big.NewFloat(2), prevZreal).Mul(currZimag, prevZimag).Add(currZimag, zImag)

		prevZreal, prevZimag = currZreal, currZimag

		// abs(z) = sqrt(a^2 + bi^2)
		squareSum := (&big.Float{}).Add(
			(&big.Float{}).Mul(currZreal, currZreal),
			(&big.Float{}).Mul(currZimag, currZimag),
		)
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func MandelbrotBigRational(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	// original z
	zReal := (&big.Rat{}).SetFloat64(float64(real(z)))
	zImag := (&big.Rat{}).SetFloat64(float64(imag(z)))

	// Nth z (previous iteration)
	prevZreal, prevZimag := &big.Rat{}, &big.Rat{}

	for n := uint8(0); n < iterations; n++ {
		// z * z = (a + bi) ^ 2 = a^2 + b^2*i^2 + 2abi = a^2 - b^2 + 2abi
		// real = a^2 - b^2
		// imag = 2abi
		currZreal := &big.Rat{}
		currZimag := &big.Rat{}

		currZreal.Mul(prevZreal, prevZreal).Sub(currZreal, (&big.Rat{}).Mul(prevZimag, prevZimag)).Add(currZreal, zReal)
		currZimag.Mul(big.NewRat(2, 1), prevZreal).Mul(currZimag, prevZimag).Add(currZimag, zImag)

		prevZreal, prevZimag = currZreal, currZimag

		// abs(z) = sqrt(a^2 + bi^2)
		squareSum := (&big.Rat{}).Add(
			(&big.Rat{}).Mul(currZreal, currZreal),
			(&big.Rat{}).Mul(currZimag, currZimag),
		)
		if squareSum.Cmp(big.NewRat(4, 1)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
