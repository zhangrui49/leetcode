package g_interview1711

import "math"

/**

有个内含单词的超大文本文件，给定任意两个不同的单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?

示例：

输入：words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
输出：1
提示：

words.length <= 100000

*/

//
func findClosest(words []string, word1 string, word2 string) int {
	array1 := []int{} // 可以用map存起来，应对多次调用优化
	array2 := []int{}
	if word1 == word2 {
		return 0
	}
	for index, v := range words {
		if v == word1 {
			array1 = append(array1, index)
			continue
		}
		if v == word2 {
			array2 = append(array2, index)
		}
	}
	i1, L1 := 0, len(array1)
	i2, L2 := 0, len(array2)
	min := math.MaxInt
	for i1 < L1 && i2 < L2 {
		distance := array2[i2] - array1[i1]
		if distance > 0 {
			i1++
		} else {
			i2++
			distance = -distance
		}
		if distance == 1 {
			return 1
		}

		if distance < min {
			min = distance
		}

	}

	return min
}

func findClosestV2(words []string, word1 string, word2 string) int {

	if word1 == word2 {
		return 0
	}
	i1 := -1
	i2 := -1
	d := math.MaxInt
	for index, v := range words {
		if v == word1 {
			i1 = index
		}
		if v == word2 {
			i2 = index
		}
		if i1 >= 0 && i2 >= 0 {
			d = min(d, abs(i2-i1))
		}

	}

	return d
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
