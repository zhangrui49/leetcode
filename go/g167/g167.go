package g167

/**

给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。

以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间

*/

//老规矩，暴力法起步(没想到暴力法通过了--)
func twoSum(numbers []int, target int) []int {

	length := len(numbers)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}

	}
	return nil
}

//把第一遍遍历的值存起来(见 g001)
func twoSumV2(numbers []int, target int) []int {

	length := len(numbers)
	values := make(map[int]int) //key:值，v:下标

	for i := 0; i < length; i++ {

		if v, ok := values[target-numbers[i]]; ok {
			return []int{v, i + 1}
		}

		values[numbers[i]] = i + 1

	}
	return nil
}

//顺序数组，解法2大材小用了，可以双指针，减少空间占用

func twoSumV3(numbers []int, target int) []int {

	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return nil
}

//顺序数组，二分查找同样可以

func twoSumV4(numbers []int, target int) []int {

	for index, _ := range numbers {
		left, right := index+1, len(numbers)-1 // 从当前位置的后面一位开始找
		for left <= right {
			mid := left + (right-left)/2
			r := target - numbers[mid] - numbers[index]
			if r > 0 {
				left = mid + 1
			} else if r < 0 {
				right = mid - 1
			} else {
				return []int{index + 1, mid + 1}
			}
		}

	}

	return nil
}
