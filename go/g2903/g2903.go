package g2903

/*
 * @lc app=leetcode.cn id=2903 lang=golang
 *
 * [2903] 找出满足差值条件的下标 I
 */

// @lc code=start
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	left := 0
	right := indexDifference
	length := len(nums)
	if right > length-1 || length == 0 {
		return []int{-1, -1}
	}
	for left < length {
		if absIntBitwise(nums[left]-nums[right]) >= valueDifference {
			return []int{left, right}
		}
		if right < length-1 {
			right++
		} else if right == length-1 {
			left++
			right = left + indexDifference
			if right > length-1 {
				break
			}
		} else {
			break
		}

	}

	return []int{-1, -1}
}

func absIntBitwise(n int) int {
	mask := n >> 63          // 获取最高位（符号位），对于负数，最高位为1，否则为0
	return (n + mask) ^ mask // 如果n是负数，mask为1，(n+1)^1会将n的符号位翻转
}

// @lc code=end
