package vs

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' {
			stack = append(stack, v)
		} else if v == ')' {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if top == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if v == '[' {
			stack = append(stack, v)
		} else if v == ']' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if v == '{' {
			stack = append(stack, v)
		} else if v == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}
	return len(stack) == 0
}

// @lc code=end
