package main

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度
*/

func lengthOfLongestSubstring(s string) int {
	hash := [256]int{}
	var left, max = 0, 0
	for i, c := range s {
		if hash[c] > left {
			left = hash[c]
		}
		l := i - left + 1
		hash[c] = i + 1
		if l > max {
			max = l
		}
	}
	return max
}

func lengthOfLongestSubstringV2(s string) int {
	cache := make(map[rune]int, 0)
	left, max := 0, 0
	for i, c := range s {
		if v, ok := cache[c]; ok {
			left = getMax(left, v+1)
		}
		cache[c] = i
		max = getMax(max, i-left+1)
	}

	return max
}

func lengthOfLongestSubstringV3(s string) int {
	cache := make(map[byte]int)
	left, max, length := 0, 0, len(s)
	for i := 0; i < length; i++ {
		c := s[i]
		if v, ok := cache[c]; ok {
			left = getMax(left, v+1)
		}
		cache[c] = i
		max = getMax(max, i-left+1)
	}

	return max
}

func getMax(x, y int) int {

	if x > y {
		return x
	}

	return y
}
