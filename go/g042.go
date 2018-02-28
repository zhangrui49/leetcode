package _go

func trap(height []int) int {
	if len(height) < 1 {
		return 0
	}
	array := getValues(height)
	array = check(array)
	return calculateTotal(array)
}

func getValues(height []int) [][]int {
	array := [][]int{}
	tempA := []int{}
	temp := height[0]
	length := len(height) - 1

	for i := 1; i < length; i++ {
		if height[i] < height[i-1] {
			temp = height[i-1]

			tempA = append(tempA, temp)
			for i < length && temp >= height[i] {
				tempA = append(tempA, height[i])
				i++
			}
		}
		if len(tempA) != 0 {
			tempA = append(tempA, height[i])
			array = append(array, tempA)
			tempA = []int{}
		}
	}
	return array
}

func check(values [][]int) [][]int {

	for i, array := range values {
		result := make([]int, len(array))
		if array[0] > array[len(array)-1] {
			for i, _ := range array {
				result[i] = array[len(array)-i-1]
			}

			values[i] = []int{0}
		}

		for _, v := range getValues(result) {
			values = append(values, v)
		}

	}

	return values
}

func calculateTotal(arrs [][]int) (total int) {

	for _, array := range arrs {

		if array[0] < array[len(array)-1] {
			for i := 1; i < len(array)-1; i++ {
				if array[0]-array[i] > 0 {
					total += array[0] - array[i]
				}

			}
		} else {
			for i := 1; i < len(array)-1; i++ {
				if array[len(array)-1]-array[i] > 0 {
					total += array[len(array)-1] - array[i]
				}

			}
		}

	}
	return
}
