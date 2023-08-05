package main

import (
	"fmt"
	"tempconv"
)

func main() {
	zeroK := tempconv.Kelvin(0.0)
	fmt.Printf("Zero Kelvin: %s\n", zeroK)
	fmt.Printf("To Celsius: %s\n", tempconv.KToC(zeroK))
	fmt.Printf("To Fahrenheit: %s\n\n", tempconv.KToF(zeroK))

	zeroC := tempconv.Celsius(0.0)
	fmt.Printf("Zero Celsius: %s\n", zeroC)
	fmt.Printf("To Fahrenheit: %s\n", tempconv.CToF(zeroC))
	fmt.Printf("To Kelvin: %s\n\n", tempconv.CToK(zeroC))

	zeroF := tempconv.Fahrenheit(0.0)
	fmt.Printf("Zero Fahrenheit: %s\n", zeroF)
	fmt.Printf("To Celsius: %s\n", tempconv.FToC(zeroF))
	fmt.Printf("To Kelvin: %s\n\n", tempconv.FToK(zeroF))

	hundredK := tempconv.Kelvin(100.0)
	fmt.Printf("100 Kelvin degrees: %s\n", hundredK)
	fmt.Printf("To Celsius: %s\n", tempconv.KToC(hundredK))
	fmt.Printf("To Fahrenheit: %s\n\n", tempconv.KToF(hundredK))

	hundredC := tempconv.Celsius(100.0)
	fmt.Printf("100 Celsius degrees: %s\n", hundredC)
	fmt.Printf("To Fahrenheit: %s\n", tempconv.CToF(hundredC))
	fmt.Printf("To Kelvin: %s\n\n", tempconv.CToK(hundredC))

	hundredF := tempconv.Fahrenheit(100.0)
	fmt.Printf("100 Fahrenheit degrees: %s\n", hundredF)
	fmt.Printf("To Celsius: %s\n", tempconv.FToC(hundredF))
	fmt.Printf("To Kelvin: %s\n\n", tempconv.FToK(hundredF))
}