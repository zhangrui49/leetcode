package g704

func search(nums []int, target int) int {
	size := len(nums)
	left, right := 0, size-1

	for left <= right {
		mid := left + (right-left)/2
		if  nums[mid] == target {
			return mid
		} else if nums[mid] < target{
			left = mid+1
		} else if nums[mid]>target {
			right = mid-1
		}
	}

	return -1
}
