package main

import (
	"strconv"
	"strings"
)

/**
规律 ：
	1*10+10 = 20
	20*10+10^2=300
	300*10+10^3=4000
*/

func countDigitOne(n int) int {
	index := 0
	for {

		index++
	}

}

func bad(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		str := itoa(i)
		count += strings.Count(str, "1")
	}
	return count
}

func itoa(val int) string {
	return strconv.Itoa(val)
}
