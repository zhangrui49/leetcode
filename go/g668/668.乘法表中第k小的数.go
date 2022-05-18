/*
 * @lc app=leetcode.cn id=668 lang=golang
 *
 * [668] 乘法表中第k小的数
 *
 * https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/description/
 *
 * algorithms
 * Hard (52.11%)
 * Likes:    264
 * Dislikes: 0
 * Total Accepted:    18.2K
 * Total Submissions: 32.3K
 * Testcase Example:  '3\n3\n5'
 *
 * 几乎每一个人都用 乘法表。但是你能在乘法表中快速找到第k小的数字吗？
 *
 * 给定高度m 、宽度n 的一张 m * n的乘法表，以及正整数k，你需要返回表中第k 小的数字。
 *
 * 例 1：
 *
 *
 * 输入: m = 3, n = 3, k = 5
 * 输出: 3
 * 解释:
 * 乘法表:
 * 1    2    3
 * 2    4    6
 * 3    6    9
 *
 * 第5小的数字是 3 (1, 2, 2, 3, 3).
 *
 *
 * 例 2：
 *
 *
 * 输入: m = 2, n = 3, k = 6
 * 输出: 6
 * 解释:
 * 乘法表:
 * 1    2    3
 * 2    4    6
 *
 * 第6小的数字是 6 (1, 2, 2, 3, 4, 6).
 *
 *
 * 注意：
 *
 *
 * m 和 n 的范围在 [1, 30000] 之间。
 * k 的范围在 [1, m * n] 之间。
 *
 *
 */
package g668

import "sort"

// @lc code=start

//必须OOM
func findKthNumber(m int, n int, k int) int {
	ints := make([]int, 0)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			ints = append(ints, i*j)
		}
	}
	sort.Ints(ints)
	return ints[k-1]
}

//二分查找，转换思想，找小于乘积X的个数，如果个数等于K，则该数字即第K个（该数字可能不在乘法表中，需要找最小的那个）
func findKthNumberV2(m int, n int, k int) int {
	left := 1
	right := m * n

	for left < right {
		mid := left + (right-left)/2 //需要找的数
		count := 0
		for i := 1; i <= m; i++ {
			mc := mid / i // x/i 为第i行小于x的数的个数（i,2i,3i,4i...）
			if mc > n {   //个数可能大于列数
				mc = n
			}
			count += mc
		}
		if count >= k { //避免该值不在乘法表中，所以等于k也将右指针左移，继续找
			right = mid
		} else {
			left = mid + 1
		}

	}
	return left
}

// @lc code=end
