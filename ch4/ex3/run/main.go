package main

import (
	"fmt"
	"os"
	"strconv"

	"reverse"
)

func main() {
	var inputsStr []string = os.Args[1:]
	var nums []int = make([]int, len(inputsStr))

	for i, e := range inputsStr {
		asint, _ := strconv.ParseInt(e, 10, 0)  // ignoring parse errors
		nums[i] = int(asint)
	}

	fmt.Printf("Original ints sequence: %v\n", nums)

	reverse.Reverse(&nums)
	fmt.Printf("Reversed sequence: %v\n\n", nums)
}
