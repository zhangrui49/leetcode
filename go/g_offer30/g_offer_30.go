package goffer30

import "math"

/**

定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)

*/

type MinStack struct {
	minStack []int
	eleStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.eleStack = append(this.eleStack, x)
	v := this.minStack[len(this.minStack)-1]
	if x <= v {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {

	v := this.Top()
	this.eleStack = this.eleStack[0 : len(this.eleStack)-1]
	min := this.minStack[len(this.minStack)-1]
	if v == min {
		this.minStack = this.minStack[0 : len(this.minStack)-1]
	}

}

func (this *MinStack) Top() int {
	return this.eleStack[len(this.eleStack)-1]
}

func (this *MinStack) Min() int {

	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
