package _go

func lengthOfLongestSubstring(s string) int {
	//hash := make([]int, 256)
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
