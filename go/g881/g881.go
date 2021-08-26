package g881

import (
	"sort"
)
//贪心算法
func numRescueBoats(people []int, limit int) int {
	//快排
	sort.Ints(people)
	start, end := 0, len(people)-1
	boats:=0
	for start <= end {
		if people[start]+people[end] > limit {
			end--
		} else {
			end--
			start++
		}
		boats++
	}
	return boats
}
