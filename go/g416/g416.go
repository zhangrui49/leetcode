package g416

/**

给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等

*/

//暴力法，找出所有可能的和，判断是否存在（可能会超内存）
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

	sums := subsets(nums)

	for _, v := range sums {
		if v == target {
			return true
		}
	}
	return false
}

func subsets(nums []int) []int {
	sum := 0
	ans := []int{}
	var dfs func(int)
	dfs = func(start int) {
		if len(nums) < start {
			return
		}

		ans = append(ans, sum)
		for i := start; i < len(nums); i++ {
			sum += nums[i]
			dfs(i + 1)
			sum -= nums[i]
		}

	}
	dfs(0)
	return ans
}

//动态规划
func canPartitionV2(nums []int) bool {
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

	dp := make([]bool, target+1)
	dp[0] = true

	for _, v := range nums {
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}
