/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode.cn/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (43.85%)
 * Likes:    675
 * Dislikes: 0
 * Total Accepted:    188.9K
 * Total Submissions: 430.5K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
 *
 * 换句话说，s1 的排列之一是 s2 的 子串 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "ab" s2 = "eidbaooo"
 * 输出：true
 * 解释：s2 包含 s1 的排列之一 ("ba").
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1= "ab" s2 = "eidboaoo"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s1.length, s2.length <= 10^4
 * s1 和 s2 仅包含小写字母
 *
 *
 */
package g567

// @lc code=start

//两个字符串A,B各个字母个数相等，则A是B的排列之一
//滑动窗口，在S2找S1长度相等字串，判断这个窗口各个字母个数与S1是否相等
func checkInclusion(s1 string, s2 string) bool {

	l1 := len(s1)
	l2 := len(s2)
	if l1 > l2 {
		return false
	}
	cc1 := make(map[byte]int)
	cc2 := make(map[byte]int)

	for i := 0; i < l1; i++ {
		c1 := s1[i]
		cc1[c1] += 1

		c2 := s2[i]
		cc2[c2] += 1
	}

	if isMapEqu(cc1, cc2) {
		return true
	}

	for i := l1; i < l2; i++ {
		left := i - l1
		lc := s2[left]
		cc2[lc]--
		if cc2[lc] == 0 {
			delete(cc2, lc)
		}
		cc2[s2[i]] += 1
		if isMapEqu(cc1, cc2) {
			return true
		}
	}

	return false
}

func isMapEqu(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if v != m2[k] {
			return false
		}
	}

	return true
}

// @lc code=end
