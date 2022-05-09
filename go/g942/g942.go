package g942

/**

由范围 [0,n] 内所有整数组成的 n + 1 个整数的排列序列可以表示为长度为 n 的字符串 s ，其中:

如果 perm[i] < perm[i + 1] ，那么 s[i] == 'I'
如果 perm[i] > perm[i + 1] ，那么 s[i] == 'D'
给定一个字符串 s ，重构排列 perm 并返回它。如果有多个有效排列perm，则返回其中 任何一个


*/

//首先暴力法(理论上会OOM)
func diStringMatch(s string) []int {
	n := len(s)
	nums := make([]int, n+1)
	for i := 0; i <= n; i++ {
		nums[i] = i
	}
	perms := permute(nums)
	for _, v := range perms {
		for i := 1; i <= n; i++ {
			var c byte
			if v[i] > v[i-1] {
				c = 'I'
			} else {
				c = 'D'
			}
			if c != s[i-1] {
				break
			}
			if i == n {
				return v
			}
		}

	}
	return nil
}

//见 g046 全排列
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	length := len(nums)
	set := make([]int, 0)
	vf := make(map[int]bool)
	for i := 0; i < length; i++ {
		vf[i] = false
	}
	var dfs func()
	dfs = func() {
		if len(set) == length {
			result = append(result, append([]int(nil), set...))
			return
		}
		for i := 0; i < length; i++ {
			if !vf[i] {
				set = append(set, nums[i])
				vf[i] = true
				dfs()
				vf[i] = false
				set = set[:len(set)-1]
			}

		}

	}
	dfs()
	return result
}

//贪心法
func diStringMatchV2(s string) []int {
	length := len(s)
	r := make([]int, length+1)
	left, right := 0, length
	for index, element := range s {
		if element == 'I' { //i+1要大于i ,所以i 取当前最小的
			r[index] = left
			left++
		} else {
			r[index] = right
			right--
		}
	}
	r[length] = left
	return r
}
