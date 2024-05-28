package g2951

/*
 * @lc app=leetcode.cn id=2951 lang=golang
 *
 * [2951] 找出峰值
 */

// @lc code=start
func findPeaks(mountain []int) []int {
	result := make([]int, 0)

	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i-1] < mountain[i] && mountain[i] > mountain[i+1] {
			result = append(result, i)
		}
	}
	return result
}

// @lc code=end
