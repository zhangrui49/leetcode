/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.70%)
 * Likes:    1671
 * Dislikes: 0
 * Total Accepted:    627.3K
 * Total Submissions: 969.4K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 *
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^4
 *
 *
 */
package g215

import (
	"sort"
)

// @lc code=start

//API大法
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

//快排整一波
func findKthLargestV2(nums []int, k int) int {

	var quickSort func(int, int)
	quickSort = func(low, high int) {
		if low >= high {
			return
		}

		left, right := low, high

		pivot := nums[(left+right)/2]

		for left <= right {
			for pivot > nums[left] {
				left++ //小于pivot的都放在左边
			}
			for pivot < nums[right] {
				right-- //大于pivot的都放在右边
			}
			//到这里，nums[left]>pivot>nums[right]了，需要swap
			if left <= right {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right--
			}
		}
		//到这里left>right了
		quickSort(low, right)
		quickSort(left, high)
	}
	quickSort(0, len(nums)-1)
	return nums[len(nums)-k]
}

//冒泡排序
func findKthLargestV3(nums []int, k int) int {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}

	}
	return nums[len(nums)-k]
}

//选择排序
func findKthLargestV4(nums []int, k int) int {

	length := len(nums)

	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 0; j < length-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		temp := nums[maxIndex]
		nums[maxIndex] = nums[length-i-1]
		nums[length-i-1] = temp
	}

	return nums[len(nums)-k]
}

//插入排序
func findKthLargestV5(nums []int, k int) int {
	length := len(nums)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
		}
	}

	return nums[len(nums)-k]
}

//堆排序
func findKthLargestV6(nums []int, k int) int {
	return nums[len(nums)-k]
}

//希尔排序(优化版插入排序)
func findKthLargestV7(nums []int, k int) int {
	length := len(nums)

	//每次拆分成n个子数组
	for n := length / 2; n > 0; n = n / 2 {
		//每个子数组插入排序
		for i := n; i < length; i++ {
			left := i
			//nums[left - n]是一组的元素
			for left-n >= 0 {
				if nums[left] < nums[left-n] {
					temp := nums[left]
					nums[left] = nums[left-n]
					nums[left-n] = temp
				}
				left -= n //步进是n
			}
		}
	}

	return nums[len(nums)-k]
}

//归并排序
func findKthLargestV8(nums []int, k int) int {
	return nums[len(nums)-k]
}

//基数排序（桶排序)(不支持负数)
func findKthLargestV9(nums []int, k int) int {
	max := nums[0]
	length := len(nums)
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	//最大值的位数决定循环次数,每次循环取更高一位
	for i := 1; max/i > 0; i = i * 10 {
		var bucket [10][]int

		for j := 0; j < length; j++ {
			v := (nums[j] / i) % 10 //取第i位的数字
			bucket[v] = append(bucket[v], nums[j])
		}

		start := 0
		for j := 0; j < 10; j++ {
			vl := len(bucket[j])
			for k := 0; k < vl; k++ {
				nums[start] = bucket[j][k]
				start++
			}
		}
	}

	return nums[len(nums)-k]
}

//基数排序（桶排序)
func findKthLargestV10(nums []int, k int) int {
	max := abs(nums[0])
	length := len(nums)
	for _, v := range nums {
		if abs(v) > max {
			max = abs(v)
		}
	}
	//最大值的位数决定循环次数,每次循环取更高一位
	for i := 1; max/i > 0; i = i * 10 {
		var bucket [10][]int
		var n_bucket [10][]int

		for j := 0; j < length; j++ {
			if nums[j] >= 0 {
				v := (nums[j] / i) % 10 //取第i位的数字
				bucket[v] = append(bucket[v], nums[j])
			} else {
				v := abs((nums[j] / i) % 10) //取第i位的数字
				n_bucket[v] = append(n_bucket[v], nums[j])
			}
		}
		//感觉有点蠢，是不是有更好的办法
		start := 0
		for j := 9; j >= 0; j-- {
			vl := len(n_bucket[j])
			for k := vl - 1; k >= 0; k-- {
				nums[start] = n_bucket[j][k]
				start++
			}
		}
		for j := 0; j < 10; j++ {
			vl := len(bucket[j])
			for k := 0; k < vl; k++ {
				nums[start] = bucket[j][k]
				start++
			}
		}
	}

	return nums[len(nums)-k]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
