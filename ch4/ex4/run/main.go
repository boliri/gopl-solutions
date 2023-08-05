package main

import (
	"fmt"
	"os"
	"strconv"

	"rotate"
)

func main() {
	var inputsStr []string = os.Args[1:]
	var nums []int = make([]int, len(inputsStr) - 2) // exclude -n arg and its value
	var positions int

	for i, e := range inputsStr {
		if e == "-n" {
			n, _ := strconv.ParseInt(inputsStr[i+1], 10, 0)  // ignoring parse errors
			positions = int(n)
			break
		}

		asint, _ := strconv.ParseInt(e, 10, 0)  // ignoring parse errors
		nums[i] = int(asint)
	}

	fmt.Printf("Original ints sequence: %v\n", nums)

	rotate.LeftRotate(nums, positions)
	fmt.Printf("Left-rotated sequence: %v\n\n", nums)
}
