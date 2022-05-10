package g016

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	dep := math.MaxInt
	sum := 0
	for i := 0; i < length; i++ {
		left, right := i+1, length-1

		for left < right {
			tempSum := nums[i] + nums[left] + nums[right]
			v := target - tempSum
			if v > 0 {
				left++
			} else if v < 0 {
				right--
			} else {
				return tempSum
			}
			if abs(v) < dep {
				dep = abs(v)
				sum = tempSum
			}

		}
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
