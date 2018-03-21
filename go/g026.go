package _go

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	start := 0

	for i := 1; i < len(nums); i++ {

		if nums[start] != nums[i] {
			start++
			nums[start] = nums[i]
		}
	}
	return start + 1
}
