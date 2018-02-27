package _go

import "fmt"

//
//func fourSum(nums []int, target int) [][]int {
//
//	result := [...][4]int{}
//	total := [...][4]int{}
//
//}

func main() {
	fmt.Print(getAll([]int{2, 7, 11, 15, 16}, 4))
}

func getAll(nums []int, per int) (slice []interface{}) {
	temp := nums

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		ret := getAll(temp, per-1)
		sigal := [4]int{}
		for k, _ := range ret {
			sigal[k] = n
		}
		slice = append(slice, sigal)
	}
	return
}
