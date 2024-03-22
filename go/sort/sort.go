package sort

func bubbleSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-i-1; j++ {
			temp := 0
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				//arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}

	}

	return arr
}

func selectSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		minPos := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minPos] {
				minPos = j
			}
		}

		arr[i], arr[minPos] = arr[minPos], arr[i]

	}

	return arr
}

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func quickSort(arr []int) []int {
	quickSortImpl(arr, 0, len(arr)-1)
	return arr
}

// 3, 1, 2, 4
func quickSortImpl(arr []int, left, right int) {
	if left > right {
		return
	}
	start := left      //0
	end := right       //3
	pivot := arr[left] //3
	for start < end {
		for pivot <= arr[end] && start < end {
			end--
		}
		for pivot >= arr[start] && start < end {
			start++
		}
		temp := arr[end]
		arr[end] = arr[start]
		arr[start] = temp
	}
	arr[left] = arr[start]
	arr[start] = pivot
	quickSortImpl(arr, left, end-1)
	quickSortImpl(arr, end+1, right)
}

func mergeSort(arr []int) []int {
	divide(arr, 0, len(arr)-1)
	return arr
}

func divide(arr []int, left, right int) []int {
	mid := (left + right) / 2
	if left < right {
		divide(arr, left, mid)
		divide(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
	return arr
}

func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i := left
	j := mid + 1
	k := 0
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			k++
			i++
		} else {
			temp[k] = arr[j]
			k++
			j++
		}
	}
	for i <= mid {
		temp[k] = arr[i]
		k++
		i++
	}
	for j <= right {
		temp[k] = arr[j]
		k++
		j++
	}
	for x := 0; x < len(temp); x++ {
		arr[x+left] = temp[x]
	}
}
