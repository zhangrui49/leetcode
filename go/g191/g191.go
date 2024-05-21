package g191

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(n int) int {

	count := 0
	for i := 0; i < 32; i++ {
		if n&(1<<i) > 0 {
			count++
		}
	}
	return count
}

// @lc code=end
