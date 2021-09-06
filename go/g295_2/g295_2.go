package g295_2

import "sort"

type MedianFinder struct {
	Data []int
}

func Constructor() MedianFinder {
	var medianFinder MedianFinder
	medianFinder.Data = make([]int, 0)
	return medianFinder
}

func (this *MedianFinder) AddNum(num int) {
	this.Data = append(this.Data, num)

}

func (this *MedianFinder) FindMedian() float64 {
	sort.Ints(this.Data)
	size := len(this.Data)

	if size&1 == 1 {
		mid := (size + 1) / 2
		result := this.Data[mid-1]
		return float64(result)
	} else {
		mid := size / 2

		result := float64(this.Data[mid-1]+this.Data[mid]) / 2.0
		return float64(result)
	}

}
