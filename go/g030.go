package main

import (
	"sort"
	"strings"
)

var result []string

func findSubstring(s string, words []string) []int {

	result = []string{}
	indexs := []int{}
	perm(words, 0)
	for _, v := range result {
		temp := s
		index := strings.Index(temp, v)
		tempIndex := 0
		for index != -1 {
			indexs = append(indexs, index+tempIndex)
			temp = string([]rune(temp)[1:])
			tempIndex += 1
			index = strings.Index(temp, v)
		}
	}
	indexs = RemoveDuplicates(indexs)
	sort.Ints(indexs)
	return indexs
}

func RemoveDuplicates(a []int) (ret []int) {
	a_len := len(a)
out:
	for i := 0; i < a_len; i++ {

		for j := i + 1; j < a_len; j++ {
			if a[i] == a[j] {
				continue out
			}
		}

		ret = append(ret, a[i])
	}
	return
}

func perm(arr []string, start int) {
	if start == len(arr) {
		con := ""
		for _, v := range arr {
			con += v
		}
		result = append(result, con)
	}

	for i := start; i < len(arr); i++ {
		swap(arr, start, i)

		perm(arr, start+1)

		swap(arr, start, i)

	}
}

func swap(arr []string, pre int, aft int) {
	s := arr[pre]
	arr[pre] = arr[aft]
	arr[aft] = s
}
