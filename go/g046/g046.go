package g046

/**

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

*/

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	length := len(nums)
	set := make([]int, 0)
	vf := make(map[int]bool) //标记该下标对应值是否用过
	for i := 0; i < length; i++ {
		vf[i] = false
	}
	var dfs func()
	dfs = func() {
		if len(set) == length {
			result = append(result, append([]int(nil), set...))
			return
		}
		for i := 0; i < length; i++ {
			if !vf[i] {
				set = append(set, nums[i])
				vf[i] = true
				dfs()
				vf[i] = false
				set = set[:len(set)-1]
			}

		}

	}
	dfs()
	return result
}
