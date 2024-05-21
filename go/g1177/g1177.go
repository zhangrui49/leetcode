package g1177

/*
 * @lc app=leetcode.cn id=1177 lang=golang
 *
 * [1177] 构建回文串检测
 */

//性能过不掉，超时
func canMakePaliQueriesBad(s string, queries [][]int) []bool {

	answers := make([]bool, len(queries))
	countMap := make(map[rune]int, 26)
	clear(countMap)
	for index, v := range queries {

		left := v[0]
		right := v[1]
		placeCount := v[2]
		for i := left; i <= right; i++ {
			char := rune(s[i])
			countMap[char] = 1 - countMap[char]
		}
		count := 0
		for key, value := range countMap {
			count += value
			countMap[key] = 0
		}
		answers[index] = count-placeCount*2 < 2
	}

	return answers
}

func clear(countMap map[rune]int) {
	for i := 'a'; i <= 'z'; i++ {
		countMap[i] = 0
	}
}

// @lc code=start
func canMakePaliQueries(s string, queries [][]int) []bool {
	parity := make([][26]int, len(s)+1)
	for i, c := range s {
		parity[i+1] = parity[i]                   //字符串s中前i个字符的统计结果
		parity[i+1][c-'a'] = 1 - parity[i][c-'a'] //c-'a'表示字符c对应的索引位置
	}
	answers := make([]bool, len(queries))

	for index, v := range queries {
		left := v[0]
		right := v[1]
		placeCount := v[2]
		count := 0
		for i := 0; i < 26; i++ {
			if parity[right+1][i]^parity[left][i] > 0 {
				count++
			}
		}
		answers[index] = count-placeCount*2 < 2
	}

	return answers
}

// @lc code=end
