package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	totalLen := len1 + len2
	flag := totalLen%2 == 1 // 奇偶
	var targetLen = 0
	if flag {
		targetLen = (totalLen + 1) / 2
	} else {
		targetLen = totalLen / 2
	}
	bigNum, left, right := 0, 0, 0
	i, j, index := 0, 0, 0

	for index < targetLen {
		if i < len1 {
			left = nums1[i]
			if j < len2 {
				right = nums2[j]
				if left >= right {
					bigNum = right
					j++
				} else {
					bigNum = left
					i++
				}
			} else {
				bigNum = nums1[i]
				i++
			}
		} else {
			bigNum = nums2[j]
			j++
		}
		index++
	}
	if flag {
		return float64(bigNum)
	} else {
		if i < len1 {
			left = nums1[i]
			if j < len2 {
				right = nums2[j]
				if left <= right {
					return float64(left+bigNum) / 2
				} else {
					return float64(right+bigNum) / 2
				}
			} else {
				return float64(nums1[i]+bigNum) / 2
			}
		} else {
			return float64(nums2[j]+bigNum) / 2
		}
	}
}
