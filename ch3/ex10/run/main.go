package main

import (
	"fmt"
	"os"

	"comma"
)

func main() {
	value := os.Args[1]
	commaValue := comma.Comma(value)

	fmt.Printf("Number %s prettified with commas is %s\n\n", value, commaValue)
}
