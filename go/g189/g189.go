package g189

/**

给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数

*/

func rotate(nums []int, k int) {

	length := len(nums)
	k = k % length //轮转超过一轮无意义
	left := make([]int, k)
	right := make([]int, length-k)
	copy(left, nums[length-k:length])
	copy(right, nums[:length-k])

	for i := 0; i < k; i++ {
		nums[i] = left[i]
	}
	for i := 0; i < length-k; i++ {
		nums[i+k] = right[i]
	}
}

func rotateV2(nums []int, k int) {

	length := len(nums)
	k = k % length //轮转超过一轮无意义
	temp := make([]int, length)
	copy(temp, nums)
	for i := 0; i < length; i++ {
		index := (i + k) % length
		nums[index] = temp[i]
	}
}

func rotateV3(nums []int, k int) {
	k = k % len(nums) //轮转超过一轮无意义
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {

	for i := 0; i < len(nums)/2; i++ {

		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]

	}
}
