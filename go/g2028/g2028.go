package g2028

/*
 * @lc app=leetcode.cn id=2028 lang=golang
 *
 * [2028] 找出缺失的观测数据
 */

// @lc code=start
func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	total := mean * (n + m)
	totalM := 0
	for _, vm := range rolls {
		totalM += vm
	}
	totalN := total - totalM

	if totalN > 6*n || totalN < n {
		return []int{}
	}

	result := make([]int, n)
	remainder := totalN % n
	if remainder == 0 {
		mean = totalN / n
		for i := 0; i < n; i++ {
			result[i] = mean
		}
		return result
	}

	meanN := totalN / n
	left := remainder
	for i := 0; i < n; i++ {
		if left > 0 {
			result[i] = meanN + 1
			left--
		} else {
			result[i] = meanN
		}
	}

	return result
}

// @lc code=end
