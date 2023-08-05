package main

import (
	"fmt"
	"os"
	"strconv"

	"conv"
)

func main() {
	value, _ := strconv.ParseFloat(os.Args[1], 64)  // Ignoring parsing errors

	celsius := conv.NewCelsius(value)
	fahrenheit := conv.NewFahrenheit(value)
	kelvin := conv.NewKelvin(value)

	pounds := conv.NewPounds(value)
	kg := conv.NewKilograms(value)

	meters := conv.NewMeters(value)
	feets := conv.NewFeets(value)

	fmt.Printf("Got value %f. Converting to multiple units...\n\n", value)

	fmt.Println("Temperature")
	fmt.Println(celsius)
	fmt.Println(fahrenheit)
	fmt.Println(kelvin)
	fmt.Println()

	fmt.Println("Mass")
	fmt.Println(pounds)
	fmt.Println(kg)
	fmt.Println()

	fmt.Println("Length")
	fmt.Println(feets)
	fmt.Println(meters)
}