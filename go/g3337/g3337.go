package g3337

/*
 * @lc app=leetcode.cn id=3337 lang=golang
 *
 * [3337] 字符串转换后的长度 II
 */

// @lc code=start

const mod = 1e9 + 7

// 矩阵乘法：a * b -> res
func multiply(a, b [26][26]int) [26][26]int {
	var res [26][26]int
	for i := 0; i < 26; i++ {
		for k := 0; k < 26; k++ {
			if a[i][k] == 0 {
				continue
			}
			for j := 0; j < 26; j++ {
				res[i][j] = (res[i][j] + a[i][k]*b[k][j]) % mod
			}
		}
	}
	return res
}

// 矩阵快速幂：计算 mat^power
func matrixPower(mat [26][26]int, power int) [26][26]int {
	// 初始化单位矩阵
	var res [26][26]int
	for i := 0; i < 26; i++ {
		res[i][i] = 1
	}
	for power > 0 {
		if power%2 == 1 {
			res = multiply(res, mat)
		}
		mat = multiply(mat, mat)
		power /= 2
	}
	return res
}

// code by doubao-1.5
func lengthAfterTransformations(s string, t int, nums []int) int {
	// 初始化字符计数
	counts := [26]int{}
	for _, v := range s {
		counts[v-'a'] = (counts[v-'a'] + 1) % mod
	}

	// 构建转移矩阵：mat[j][d] 表示字符j经过1次转换后生成d的数量
	var mat [26][26]int
	for j := 0; j < 26; j++ {
		move := nums[j]
		if move == 0 {
			mat[j][j] = 1 // 移动0次时保留原字符
			continue
		}
		// 计算每个j生成的d的数量
		circle := move / 26
		r := move % 26
		// 完整周期贡献：每个周期生成26个字符（j+1到j+26）
		addCircle := circle % mod
		for d := 0; d < 26; d++ {
			pos := (j + 1 + d) % 26
			mat[j][pos] = (mat[j][pos] + addCircle) % mod
		}
		// 剩余次数贡献：生成j+1到j+r的字符
		for k := 1; k <= r; k++ {
			pos := (j + k) % 26
			mat[j][pos] = (mat[j][pos] + 1) % mod
		}
	}

	// 计算转移矩阵的t次幂
	powerMat := matrixPower(mat, t)

	// 初始计数与幂矩阵相乘，得到最终计数
	finalCounts := [26]int{}
	for j := 0; j < 26; j++ {
		if counts[j] == 0 {
			continue
		}
		for d := 0; d < 26; d++ {
			finalCounts[d] = (finalCounts[d] + counts[j]*powerMat[j][d]) % mod
		}
	}

	// 计算总长度
	total := 0
	for _, v := range finalCounts {
		total = (total + v) % mod
	}
	return total
}

// 跑到 529/536 超时了
func lengthAfterTransformationsTimeLimit(s string, t int, nums []int) int {
	total := 0
	mod := int(1e9 + 7)
	counts := make([]int, 26)
	for _, v := range s {
		counts[int(v-'a')]++
	}

	for i := 0; i < t; i++ {
		newCounts := make([]int, 26)
		for j := 0; j < 26; j++ {
			count := counts[j]
			if count == 0 {
				continue
			}
			move := nums[j]
			if move == 0 {
				continue
			}
			circle := move / 26
			if circle > 0 {
				add := circle * count % mod
				for k := 1; k <= 26; k++ {
					d := (j + k) % 26
					newCounts[d] = (newCounts[d] + add) % mod
				}
			}
			r := move % 26
			if r == 0 {
				continue
			}
			for k := 1; k <= r; k++ {
				d := (j + k) % 26
				//这里提前取模，避免数字越界
				newCounts[d] = (count + newCounts[d]) % mod
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
