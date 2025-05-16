package g3337

/*
 * @lc app=leetcode.cn id=3337 lang=golang
 *
 * [3337] 字符串转换后的长度 II
 */

// @lc code=start
func lengthAfterTransformations(s string, t int, nums []int) int {
	total := 0
	trans := make([]int, len(s))
	for i, v := range s {
		trans[i] = int(v - 'a')
	}

	for i := 0; i < t; i++ {
		result := make([]int, 0)
		for j := 0; j < len(trans); j++ {
			origin := trans[j]
			move := nums[origin]
			for k := 1; k <= move; k++ {
				result = append(result, (origin+k)%26)
			}
		}
		trans = result
	}
	total = len(trans)
	total = total % (1e9 + 7)
	return total
}

// @lc code=end
