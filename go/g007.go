package main

import (
	"math"
)

func Reverse(x int) int {
	var result = 0
	var num = 0
	for x != 0 {
		num = x % 10
		result = num + result*10
		x = x / 10
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
