package g295

//数据流的中位数
type MedianFinder struct {
	Median *ListNode
	Count  int
}

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

func Constructor() MedianFinder {
	var medianFinder MedianFinder
	medianFinder.Count = 0
	return medianFinder
}

func (listNode *ListNode) appendNext(list *ListNode) {
	next := listNode.Next
	if next == nil {
		listNode.Next = list
	} else {
		listNode.Next = list
		list.Prev = listNode
		list.Next = next
	}
}

func (listNode *ListNode) appendPrev(list *ListNode) {
	prev := listNode.Next
	if prev == nil {
		listNode.Prev = list
	} else {
		listNode.Prev = list
		list.Prev = prev
		list.Next = listNode
	}
}

func (this *MedianFinder) AddNum(num int) {
	var list ListNode
	list.Val = num
	this.Count++
	if this.Count == 1 {
		this.Median = &list
		return
	}

	if num >= this.Median.Val {
		next := this.Median.Next
		if next == nil {
			this.Median.appendNext(&list)
		} else {
			for next != nil {
				if next.Val <= num {
					next = next.Next
					continue
				}
				break
			}
			next.appendPrev(&list)
		}

		this.Median = this.Median.Next

	} else {

		prev := this.Median.Prev
		if prev == nil {
			this.Median.appendPrev(&list)
		} else {
			for prev != nil {
				if prev.Val >= num {
					prev = prev.Prev
					continue
				}
				break
			}
			prev.appendNext(&list)
		}
		this.Median = this.Median.Prev
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.Count&1 == 1 {
		return float64(this.Median.Val)
	} else {

		return float64((this.Median.Val + this.Median.Next.Val)) / 2.0
	}
}
