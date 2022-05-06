package g053

import "math"

/**

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分

*/

//贪心法
func maxSubArray(nums []int) int {

	max, i, sum := math.MinInt32, 0, 0
	length := len(nums)

	for i < length {
		sum = sum + nums[i]
		max = getMax(max, sum)
		if sum < 0 {
			sum = 0
		}
		i++
	}

	return max

}

//暴力法
func maxSubArrayV2(nums []int) int {

	max, sum := nums[0], 0
	length := len(nums)

	for i := 0; i < length; i++ {
		sum = 0
		for j := i; j < length; j++ {
			sum = sum + nums[j]
			if sum > max {
				max = sum
			}
		}
	}

	return max

}

//动态规划
func maxSubArrayV3(nums []int) int {

	max := nums[0]
	length := len(nums)
	sums := make([]int, length)
	sums[0] = nums[0]

	for i := 1; i < length; i++ {
		sums[i] = getMax(sums[i-1]+nums[i], nums[i])
		max = getMax(max, sums[i])
	}

	return max

}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
