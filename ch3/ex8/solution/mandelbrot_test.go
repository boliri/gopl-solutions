package mandelbrot

import (
	"testing"
)


func BenchmarkMandelbrotComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MandelbrotComplex64(complex128(complex(float64(i), float64(i))))
	}
}

func BenchmarkMandelbrotComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MandelbrotComplex128(complex128(complex(float64(i), float64(i))))
	}
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MandelbrotBigFloat(complex128(complex(float64(i), float64(i))))
	}
}

func BenchmarkMandelbrotBigRational(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MandelbrotBigRational(complex128(complex(float64(i), float64(i))))
	}
}