//数据流的中位数
package g295_2

import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	max *MaxHeap
	min *MinHeap
}

func Constructor() MedianFinder {
	var medianFinder MedianFinder
	medianFinder.max = &MaxHeap{}
	medianFinder.min = &MinHeap{}
	return medianFinder
}

func (this *MedianFinder) AddNum(num int) {

	if this.min.Len() == 0 || num <= this.min.IntSlice[0] {
		heap.Push(this.min, num)
		if this.max.Len()+1 < this.min.Len() {
			heap.Push(this.max, heap.Pop(this.min))
		}

	} else {

		heap.Push(this.max, num)

		if this.max.Len() > this.min.Len() {
			heap.Push(this.min, heap.Pop(this.max))
		}

	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.min.Len() == 0 {
		return 0
	}

	if this.max.Len() == this.min.Len() {
		return float64(this.min.IntSlice[0]+this.max.IntSlice[0]) / 2.0
	}
	return float64(this.min.IntSlice[0])
}

type MaxHeap struct {
	MedianHeap
}

type MinHeap struct {
	MedianHeap
}

type MedianHeap struct {
	sort.IntSlice
}

func (mh *MedianHeap) Push(v interface{}) {
	mh.IntSlice = append(mh.IntSlice, v.(int))
}

func (mh *MedianHeap) Pop() interface{} {
	temp := mh.IntSlice
	result := temp[len(temp)-1]
	mh.IntSlice = temp[:len(temp)-1]
	return result
}

func (max MaxHeap) Less(i, j int) bool {
	return max.IntSlice[i] < max.IntSlice[j]
}
func (max MinHeap) Less(i, j int) bool {
	return max.IntSlice[i] > max.IntSlice[j]
}
