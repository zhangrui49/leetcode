//数据流的中位数
package g295

type MedianFinder struct {
	Median *MyListNode
	Count  int
	Direc  int
}

type MyListNode struct {
	Val  int
	Next *MyListNode
	Prev *MyListNode
}

func Constructor() MedianFinder {
	var medianFinder MedianFinder
	medianFinder.Count = 0
	return medianFinder
}

func (MyListNode *MyListNode) appendNext(list *MyListNode) {
	next := MyListNode.Next
	if next == nil {
		MyListNode.Next = list
		list.Prev = MyListNode
	} else {
		MyListNode.Next = list
		list.Prev = MyListNode
		list.Next = next
		next.Prev = list
	}
}

func (MyListNode *MyListNode) appendPrev(list *MyListNode) {
	prev := MyListNode.Prev
	if prev == nil {
		MyListNode.Prev = list
		list.Next = MyListNode
	} else {
		MyListNode.Prev = list
		list.Next = MyListNode

		list.Prev = prev
		prev.Next = list
	}
}

func (this *MedianFinder) AddNum(num int) {
	var list = new(MyListNode)
	list.Val = num
	this.Count++
	if this.Count == 1 {
		this.Median = list
		return
	}

	if num >= this.Median.Val {
		next := this.Median.Next
		if next == nil {
			this.Median.appendNext(list)
		} else {
			for next != nil {
				if next.Val <= num {
					tempNext := next.Next
					if tempNext == nil {
						break
					}
					next = tempNext
					continue
				}
				break
			}
			if next.Val <= list.Val {
				next.appendNext(list)
			} else {
				next.appendPrev(list)
			}

		}
		if this.Count&1 == 0 || this.Direc == 1 {
			this.Median = this.Median.Next
		}
		this.Direc = 0
	} else {

		prev := this.Median.Prev
		if prev == nil {
			this.Median.appendPrev(list)
		} else {
			for prev != nil {
				if prev.Val >= num {
					tempPrev := prev.Prev
					if tempPrev == nil {
						break
					}
					prev = tempPrev
					continue
				}
				break
			}
			if prev.Val <= list.Val {
				prev.appendNext(list)
			} else {
				prev.appendPrev(list)
			}

		}
		if this.Count&1 == 0 || this.Direc == 0 {
			this.Median = this.Median.Prev
		}
		this.Direc = 1
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.Count&1 == 1 {

		return float64(this.Median.Val)

	} else {
		if this.Direc == 1 {
			return float64((this.Median.Val + this.Median.Next.Val)) / 2.0
		} else {
			return float64((this.Median.Val + this.Median.Prev.Val)) / 2.0
		}

	}
}
