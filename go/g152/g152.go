package g152

/**

给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 32-位 整数。

子数组 是数组的连续子序列

*/

func maxProduct(nums []int) int {

	length := len(nums)
	max := nums[0]
	mul := nums[0]
	for i := 0; i < length; i++ {

		mul = nums[i]
		if mul > max {
			max = mul
		}
		for j := i + 1; j < length; j++ {
			mul = mul * nums[j]
			if mul > max {
				max = mul
			}

		}

	}

	return max

}

func maxProductV2(nums []int) int {

	length := len(nums)
	max := nums[0]
	mul := 1
	left, right := 0, 0

	for right < length {
		mul = mul * nums[right]
		if mul > max {
			max = mul
		}
		for max > mul && left < length {
			mul = mul / nums[left]
			left++
		}

		right++
	}

	return max

}
