// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package conv has the ability to convert plain float64 numbers to multiple units (temperature, length, mass, etc.)
package conv

import "fmt"

type Printable interface {
	String() string
}

type Metric struct {
	value    float64
	unit     string
}

func (m Metric) String() string {
	return fmt.Sprintf("%f %s", m.value, m.unit)
}

//!-
