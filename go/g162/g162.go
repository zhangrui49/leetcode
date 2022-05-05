/*
峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组 nums，找到峰值元素并返回其索引。
数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题

*/
package g162

//寻找峰值

func findPeakElement(nums []int) int {

	size := len(nums)

	if size == 1 {
		return 0
	}

	for i := 1; i < size-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}

	}

	//边界
	if 

	return -1
}
