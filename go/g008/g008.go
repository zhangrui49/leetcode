package g008

import (
	"math"
	"strings"
)

func Atoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	sign := 1
	result := 0
	if str[0] == '-' {
		sign = -1
		str = str[1:]
	} else if str[0] == '+' {
		sign = 1
		str = str[1:]
	}
	for _, v := range str {
		if v < '0' || v > '9' {
			break
		}
		result = result*10 + int(v-'0')
		if result > math.MaxInt32 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
	}
	return result * sign
}
