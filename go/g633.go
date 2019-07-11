package main

import (
	"fmt"
	"math"
)

func JudgeSquareSum(sum int) bool {
	i := 0
	j := int(math.Sqrt(float64(sum)))

	for {

		v := i*i + j*j
		if v == sum {
			fmt.Println(i, j)
			return true
		}
		if v < sum {
			i++
		} else {
			j--
		}

		if i > j {
			break
		}
	}
	return false

}
