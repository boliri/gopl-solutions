package maxmin

import (
	"fmt"
	"math"
)

func Max(nums ...int64) int64 {
	var res int64
	for _, n := range nums {
		if n > res {
			res = n
		}
	}
	return res
}

func Min(nums ...int64) int64 {
	var res int64 = math.MaxInt64
	for _, n := range nums {
		if n < res {
			res = n
		}
	}
	return res
}

func MaxAtLeastOne(nums ...int64) (int64, error) {
	if len(nums) == 0 {
		return int64(math.NaN()), fmt.Errorf("no arguments provided")
	}

	var res int64
	for _, n := range nums {
		if n > res {
			res = n
		}
	}
	return res, nil
}

func MinAtLeastOne(nums ...int64) (int64, error) {
	if len(nums) == 0 {
		return int64(math.NaN()), fmt.Errorf("no arguments provided")
	}

	var res int64 = math.MaxInt64
	for _, n := range nums {
		if n < res {
			res = n
		}
	}
	return res, nil
}
