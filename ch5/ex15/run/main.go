package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"maxmin"
)

func main() {
	var fn func(...int64) (int64, error)

	op := os.Args[1]
	switch op{
	case "max":
		fn = maxmin.MaxAtLeastOne
	case "min":
		fn = maxmin.MinAtLeastOne
	default:
		fmt.Printf("unsupported operation: %s\n", op)
		os.Exit(1)
	}

	var nums []int64
	for _, n := range os.Args[2:] {
		p, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			fmt.Printf("invalid argument: %s\n", n)
			os.Exit(1)
		}

		nums = append(nums, p)
	}

	res, err := fn(nums...)
	if err != nil {
		fmt.Printf("could not calculate value: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s(%s) = %d\n", op, strings.Join(os.Args[2:], ", "), res)
	}
}