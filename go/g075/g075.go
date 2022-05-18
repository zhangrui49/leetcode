package g075

/**
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库的sort函数的情况下解决这个问题

*/

func sortColors(nums []int) {
	c0, c1, c2 := 0, 0, 0
	for _, element := range nums {
		if element == 0 {
			c0++
		} else if element == 1 {
			c1++
		} else if element == 2 {
			c2++
		}
	}
	length := len(nums)
	for i := 0; i < length; i++ {
		if i < c0 {
			nums[i] = 0
		} else if c0 <= i && i < c0+c1 {
			nums[i] = 1
		} else {
			nums[i] = 2
		}

	}

}
