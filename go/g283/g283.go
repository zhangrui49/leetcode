package g283

/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作
*/

func moveZeroes(nums []int) {

	length := len(nums)
	left, right := 0, length-1
	for left < right {
		if nums[left] == 0 {
			move(nums, left, right)
			right--
		} else {
			left++
		}
	}

}

func moveZeroesV2(nums []int) {

	length := len(nums)
	left, right := 0, 0
	for right < length {
		if nums[right] != 0 {
			if nums[left] == 0 {
				swap(nums, left, right)
			}

			left++
		}
		right++
	}

}

func move(nums []int, src, dest int) {
	temp := nums[src]
	for i := src; i < dest; i++ {
		nums[i] = nums[i+1]
	}
	nums[dest] = temp
}

func swap(nums []int, src, dest int) {
	nums[src], nums[dest] = nums[dest], nums[src]
}
