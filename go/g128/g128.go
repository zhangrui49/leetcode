package main

import "math"

func longestConsecutive(nums []int) int {
	usedMap := make(map[int]bool)
	result := 0
	for _, v := range nums {
		usedMap[v] = false
	}
	for _, e := range nums {
		used, ok := usedMap[e]
		if ok && used {
			continue
		}
		left := e - 1
		usedL, okl := usedMap[left]
		for okl && !usedL && left > math.MinInt32 {
			usedMap[left] = true
			left--
			usedL, okl = usedMap[left]
		}
		right := e + 1
		usedR, okr := usedMap[right]
		for okr && !usedR && right < math.MaxInt32 {
			usedMap[right] = true
			right++
			usedR, okr = usedMap[right]
		}
		if right-left-1 > result {
			result = right - left - 1
		}
	}
	return result
}
