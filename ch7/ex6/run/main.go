package main

import (
	"flag"
	"fmt"

	"tempflag"
)

var temp = tempflag.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
