package g3337

/*
 * @lc app=leetcode.cn id=3337 lang=golang
 *
 * [3337] 字符串转换后的长度 II
 */

// @lc code=start
// 跑到 529/536 超时了

func lengthAfterTransformations(s string, t int, nums []int) int {
	total := 0
	mod := int(1e9 + 7)
	counts := make([]int, 26)
	for _, v := range s {
		counts[int(v-'a')]++
	}

	for i := 0; i < t; i++ {
		newCounts := make([]int, 26)
		for j := 0; j < 26; j++ {
			if counts[j] == 0 {
				continue
			}
			origin := counts[j]
			move := nums[j]
			for k := 1; k <= move; k++ {
				d := (j + k) % 26
				//这里提前取模，避免数字越界
				newCounts[d] = (origin + newCounts[d]) % mod
			}
		}
		counts = newCounts
	}
	for _, v := range counts {
		total = (total + v) % mod
	}
	return total
}

// 通过计数优化内存，但是数字会越界
func lengthAfterTransformationsOverSize(s string, t int, nums []int) int {
	total := 0
	counts := make([]int, 26)
	for _, v := range s {
		counts[int(v-'a')]++
	}

	for i := 0; i < t; i++ {
		newCounts := make([]int, 26)
		for j := 0; j < 26; j++ {
			if counts[j] == 0 {
				continue
			}
			origin := counts[j]
			move := nums[j]
			for k := 1; k <= move; k++ {
				d := (j + k) % 26
				newCounts[d] += origin
			}
		}
		counts = newCounts
	}
	for _, v := range counts {
		total += v
	}
	total = total % (1e9 + 7)
	return total
}

// 暴力法，内存会超标
func lengthAfterTransformationsLowMem(s string, t int, nums []int) int {
	total := 0
	trans := make([]int, len(s))
	for i, v := range s {
		trans[i] = int(v - 'a')
	}

	for i := 0; i < t; i++ {
		result := make([]int, 0)
		for j := 0; j < len(trans); j++ {
			origin := trans[j]
			move := nums[origin]
			for k := 1; k <= move; k++ {
				result = append(result, (origin+k)%26)
			}
		}
		trans = result
	}
	total = len(trans)
	total = total % (1e9 + 7)
	return total
}

// @lc code=end
