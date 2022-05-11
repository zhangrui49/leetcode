package main

import "sort"

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	result := [][]int{}
	for i, v := range nums {
		if i > 0 && nums[i-1] == v {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[l] + nums[i] + nums[r]

			switch {
			case s < 0:
				l++
			case s > 0:
				r--
			default:
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r {
					if nums[l] == nums[l+1] {
						l++
					} else if nums[r] == nums[r-1] {
						r--
					} else {
						r--
						l++
						break
					}
				}
			}
		}
	}
	return result
}
