package g2981

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=2981 lang=golang
 *
 * [2981] 找出出现至少三次的最长特殊子字符串 I
 */

// @lc code=start
func maximumLength(s string) int {
	length := len(s)
	if length == 0 {
		return -1
	}
	left := 0
	charArray := [26][]int{}
	for left < length {
		origin := left
		right := origin + 1
		char := s[left] - 'a'
		for right < length && s[left] == s[right] {
			right++
		}

		fmt.Printf(" %c %d \n", s[left], right-origin)
		charArray[char] = append(charArray[char], right-origin)
		left = right
	}
	result := -1
	//fmt.Printf(" %v ", charArray)
	for _, array := range charArray {
		arrayLen := len(array)
		if arrayLen == 0 {
			continue
		}
		sort.Slice(array, func(i, j int) bool {
			return array[i] > array[j]
		})
		max := array[0]

		if max < 3 && arrayLen == 1 {
			continue
		}
		if arrayLen == 2 && max == 1 {
			continue
		}
		var innderMax int
		if max-2 > 1 {
			innderMax = max - 2
		} else {
			innderMax = 1
		}
		if arrayLen >= 3 && max == array[1] && max == array[2] {
			innderMax = max
		} else if arrayLen >= 2 && max-array[1] <= 1 {
			innderMax = max - 1
		}
		if innderMax > result {
			result = innderMax
		}
	}

	return result
}

// @lc code=end
