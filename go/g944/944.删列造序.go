/*
 * @lc app=leetcode.cn id=944 lang=golang
 *
 * [944] 删列造序
 *
 * https://leetcode.cn/problems/delete-columns-to-make-sorted/description/
 *
 * algorithms
 * Easy (68.37%)
 * Likes:    80
 * Dislikes: 0
 * Total Accepted:    41.2K
 * Total Submissions: 59.2K
 * Testcase Example:  '["cba","daf","ghi"]'
 *
 * 给你由 n 个小写字母字符串组成的数组 strs，其中每个字符串长度相等。
 *
 * 这些字符串可以每个一行，排成一个网格。例如，strs = ["abc", "bce", "cae"] 可以排列为：
 *
 *
 * abc
 * bce
 * cae
 *
 * 你需要找出并删除 不是按字典序升序排列的 列。在上面的例子（下标从 0 开始）中，列 0（'a', 'b', 'c'）和列 2（'c', 'e',
 * 'e'）都是按升序排列的，而列 1（'b', 'c', 'a'）不是，所以要删除列 1 。
 *
 * 返回你需要删除的列数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：strs = ["cba","daf","ghi"]
 * 输出：1
 * 解释：网格示意如下：
 * ⁠ cba
 * ⁠ daf
 * ⁠ ghi
 * 列 0 和列 2 按升序排列，但列 1 不是，所以只需要删除列 1 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：strs = ["a","b"]
 * 输出：0
 * 解释：网格示意如下：
 * ⁠ a
 * ⁠ b
 * 只有列 0 这一列，且已经按升序排列，所以不用删除任何列。
 *
 *
 * 示例 3：
 *
 *
 * 输入：strs = ["zyx","wvu","tsr"]
 * 输出：3
 * 解释：网格示意如下：
 * ⁠ zyx
 * ⁠ wvu
 * ⁠ tsr
 * 所有 3 列都是非升序排列的，所以都要删除。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == strs.length
 * 1
 * 1
 * strs[i] 由小写英文字母组成
 *
 *
 */
package g944

// @lc code=start
func minDeletionSize(strs []string) int {
	rows := len(strs)
	if rows <= 1 {
		return 0
	}
	count := 0
	cols := len(strs[0])

	for i := 0; i < cols; i++ {

		for j := 1; j < rows; j++ {
			if strs[j][i] < strs[j-1][i] {
				count++
				break
			}
		}

	}

	return count
}

// @lc code=end
