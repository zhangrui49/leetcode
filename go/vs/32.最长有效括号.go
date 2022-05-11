package vs

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode.cn/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (36.41%)
 * Likes:    1819
 * Dislikes: 0
 * Total Accepted:    265.8K
 * Total Submissions: 729.8K
 * Testcase Example:  '"(()"'
 *
 * 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "(()"
 * 输出：2
 * 解释：最长有效括号子串是 "()"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = ")()())"
 * 输出：4
 * 解释：最长有效括号子串是 "()()"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = ""
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s[i] 为 '(' 或 ')'
 *
 *
 *
 *
 */

// @lc code=start
func longestValidParentheses(s string) int {
	max := 0
	stack := []int{-1}
	length := len(s)

	for i := 0; i < length; i++ {
		c := s[i]
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				//如果长度为0，则说明）比（多一个，从这里开始不是有效的括号，记录当前位置
				stack = append(stack, i) //栈底存最后一个没有被匹配的右括号的下标
			} else {
				//长度不为0
				top := stack[len(stack)-1] //栈顶的（的下标
				vl := i - top              //匹配的括号长度
				if vl > max {
					max = vl
				}

			}

		}

	}

	return max
}

// @lc code=end
