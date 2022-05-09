package g977

/**

给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

*/

//不乱序，平方的最大值肯定出现在数组的两端
func sortedSquares(nums []int) []int {
	length := len(nums)
	left, right, index := 0, length-1, length-1
	r := make([]int, length)

	for index >= 0 {
		leftM := nums[left] * nums[left]
		rightM := nums[right] * nums[right]
		if rightM > leftM {
			r[index] = rightM
			right--
		} else {
			r[index] = leftM
			left++
		}
		index--
	}
	return r
}
