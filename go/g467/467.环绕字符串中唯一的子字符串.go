/*
 * @lc app=leetcode.cn id=467 lang=golang
 *
 * [467] 环绕字符串中唯一的子字符串
 *
 * https://leetcode.cn/problems/unique-substrings-in-wraparound-string/description/
 *
 * algorithms
 * Medium (44.81%)
 * Likes:    223
 * Dislikes: 0
 * Total Accepted:    13.6K
 * Total Submissions: 28.9K
 * Testcase Example:  '"a"'
 *
 * 把字符串 s 看作是 “abcdefghijklmnopqrstuvwxyz” 的无限环绕字符串，所以 s 看起来是这样的：
 *
 *
 * "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd...." .
 *
 *
 * 现在给定另一个字符串 p 。返回 s 中 唯一 的 p 的 非空子串 的数量 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: p = "a"
 * 输出: 1
 * 解释: 字符串 s 中只有一个"a"子字符。
 *
 *
 * 示例 2:
 *
 *
 * 输入: p = "cac"
 * 输出: 2
 * 解释: 字符串 s 中的字符串“cac”只有两个子串“a”、“c”。.
 *
 *
 * 示例 3:
 *
 *
 * 输入: p = "zab"
 * 输出: 6
 * 解释: 在字符串 s 中有六个子串“z”、“a”、“b”、“za”、“ab”、“zab”。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= p.length <= 10^5
 * p 由小写英文字母构成
 *
 *
 */

package g467

//核心是两个结尾或者开头相同的连续字符串，短的必定是长的字串
// @lc code=start
func findSubstringInWraproundString(p string) int {
	count, length := 0, len(p)
	array := [26]int{}

	for i, j := 0, 1; j <= length; j++ {
		if j == length || (p[j]-p[j-1]+26)%26 != 1 {
			for ; i < j; i++ {
				array[p[i]-'a'] = max(array[p[i]-'a'], j-i)
			}
		}
	}

	for i := 0; i < 26; i++ {
		count += array[i]
	}

	return count
}
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end
