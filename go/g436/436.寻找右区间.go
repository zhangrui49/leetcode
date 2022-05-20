/*
 * @lc app=leetcode.cn id=436 lang=golang
 *
 * [436] 寻找右区间
 *
 * https://leetcode.cn/problems/find-right-interval/description/
 *
 * algorithms
 * Medium (49.45%)
 * Likes:    137
 * Dislikes: 0
 * Total Accepted:    19.2K
 * Total Submissions: 35.6K
 * Testcase Example:  '[[1,2]]'
 *
 * 给你一个区间数组 intervals ，其中 intervals[i] = [starti, endi] ，且每个 starti 都 不同 。
 *
 * 区间 i 的 右侧区间 可以记作区间 j ，并满足 startj >= endi ，且 startj 最小化 。
 *
 * 返回一个由每个区间 i 的 右侧区间 的最小起始位置组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,2]]
 * 输出：[-1]
 * 解释：集合中只有一个区间，所以输出-1。
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[3,4],[2,3],[1,2]]
 * 输出：[-1,0,1]
 * 解释：对于 [3,4] ，没有满足条件的“右侧”区间。
 * 对于 [2,3] ，区间[3,4]具有最小的“右”起点;
 * 对于 [1,2] ，区间[2,3]具有最小的“右”起点。
 *
 *
 * 示例 3：
 *
 *
 * 输入：intervals = [[1,4],[2,3],[3,4]]
 * 输出：[-1,2,-1]
 * 解释：对于区间 [1,4] 和 [3,4] ，没有满足条件的“右侧”区间。
 * 对于 [2,3] ，区间 [3,4] 有最小的“右”起点。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 2 * 10^4
 * intervals[i].length == 2
 * -10^6 <= starti <= endi <= 10^6
 * 每个间隔的起点都 不相同
 *
 *
 */
package g436

import (
	"math"
	"sort"
)

// @lc code=start
//老规矩，暴力法起步
func findRightInterval(intervals [][]int) []int {
	length := len(intervals)

	r := make([]int, length)

	for i := 0; i < length; i++ {
		array := intervals[i]

		min := math.MaxInt
		for j := 0; j < length; j++ {
			//右侧区间能在左边就算了，还能是自己
			// if j == i {
			// 	continue
			// }
			rightArray := intervals[j]

			if rightArray[0] >= array[1] {
				if rightArray[0] < min {
					min = rightArray[0]
					r[i] = j
				}
			}
		}
		if min == math.MaxInt {
			r[i] = -1
		}
	}
	return r
}

//对startj排序，然后二分查找
func findRightIntervalV2(intervals [][]int) []int {
	length := len(intervals)
	r := make([]int, length)
	//将下标存在第三位
	for index, _ := range intervals {
		intervals[index] = append(intervals[index], index)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] //根据第一个元素升序排列
	})

	for i := 0; i < length; i++ {
		array := intervals[i]
		index := sort.Search(length, func(i int) bool { //二分查找
			return intervals[i][0] >= array[1]
		})
		if index < length {
			r[array[2]] = intervals[index][2]
		} else {
			r[array[2]] = -1
		}
	}
	return r
}

func findRightIntervalV3(intervals [][]int) []int {
	length := len(intervals)
	r := make([]int, length)
	type element struct {
		value    int
		position int
	}
	startSlice := make([]element, length)
	endSlice := make([]element, length)
	for i, v := range intervals {
		startSlice[i] = element{
			value:    v[0],
			position: i,
		}
		endSlice[i] = element{
			value:    v[1],
			position: i,
		}
	}
	sort.Slice(startSlice, func(i, j int) bool {
		return startSlice[i].value < startSlice[j].value
	})
	sort.Slice(endSlice, func(i, j int) bool {
		return endSlice[i].value < endSlice[j].value
	})
	left := 0
	for _, v := range endSlice {

		for left < length && startSlice[left].value < v.value {

			left++
		}
		if left < length {
			r[v.position] = startSlice[left].position
		} else {
			r[v.position] = -1
		}
	}

	return r
}

// @lc code=end
