package main

import (
	"strconv"
	"strings"
)

func NearestPalindromic(n string) string {
	intV, err := strconv.Atoi(n)
	// if isPalindrome(intV) {
	// 	return n
	// }
	if err != nil {
		return ""
	}
	i, j := 1, 1
	for {
		next := intV + i
		if j <= intV {
			prev := intV - j
			if isPalindrome(prev) {
				return strconv.Itoa(prev)
			}
		}
		if isPalindrome(next) {
			return strconv.Itoa(next)
		}
		i++
		j++
	}
}

func isPalindrome(number int) bool {
	str := strconv.Itoa(number)
	strV := reverse(str)
	if strings.Compare(str, strV) == 0 {
		return true
	}
	return false
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
