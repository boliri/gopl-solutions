package main

import (
	"fmt"
	"os"
	"strconv"

	"popcount"
)

func main() {
	value, _ := strconv.ParseUint(os.Args[1], 10, 64)  // Ignoring parsing errors
	population := popcount.PopCountClearRightmostNonZeroBit(value)

	fmt.Printf("Population count for %d is %d\n\n", value, population)
}