package g416

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum, max, target := 0, 0, 0

	for _, num := range nums {
		if max < num {
			max = num
		}
		sum += num
	}
	if sum&1 == 1 {
		//和不能为奇数
		return false
	}
	target = sum / 2

	if target < max {
		//最大值大于一半,则无法对半分
		return false
	}

	return true
}
