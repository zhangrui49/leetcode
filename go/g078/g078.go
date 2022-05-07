package g078

/**

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集

*/

//迭代法 000-111共2的三次方种组合
func subsets(nums []int) [][]int {

	length := len(nums)
	r := [][]int{}
	for m := 0; m < (1 << length); m++ {
		sub := []int{}

		for i, v := range nums {
			if m>>i&1 > 0 { //对应位是1，则取该下标的值
				sub = append(sub, v)
			}
		}
		r = append(r, append([]int{}, sub...))
	}
	return r
}

//递归
func subsetsV2(nums []int) [][]int {
	set := []int{}
	ans := [][]int{}
	var dfs func(int)
	dfs = func(start int) {
		if len(nums) < start {
			return
		}

		ans = append(ans, append([]int(nil), set...))
		for i := start; i < len(nums); i++ {
			set = append(set, nums[i])
			dfs(i + 1)
			set = set[:len(set)-1]
		}

	}
	dfs(0)
	return ans
}
