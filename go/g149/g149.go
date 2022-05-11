package main

import "math"

func maxPoints(points [][]int) int {

	length := len(points)
	if length <= 1 {
		return length
	}
	max := 0
	for i := 0; i < length-2; i++ {
		same := 1
		for j := i + 1; j < length; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				same++
				continue
			}
			count := same + 1
			for k := j + 1; k < length; k++ {
				if (points[j][1]-points[i][1])*(points[k][0]-points[i][0]) == (points[k][1]-points[i][1])*(points[j][0]-points[i][0]) {
					count++
				}
			}
			max = int(math.Max(float64(max), math.Max(float64(same), float64(count))))
		}
		if max <= 0 {
			max = same
		}
	}
	return max
}
