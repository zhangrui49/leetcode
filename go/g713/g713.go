package g713

/*
给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		origin := nums[i]

		if origin >= k {
			continue
		} else {
			count += 1
		}
		if i == length-1 {
			break
		}

		for j := i + 1; j < length; j++ {
			origin = origin * nums[j]

			if origin >= k {
				break
			} else {
				count += 1
			}
		}

	}

	return count
}
