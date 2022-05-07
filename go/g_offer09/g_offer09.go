package goffer09

/**

用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

*/

type CQueue struct {
	StackIn  []int
	StackOut []int
}

func Constructor() CQueue {
	return CQueue{
		StackIn:  make([]int, 0),
		StackOut: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.StackIn = append(this.StackIn, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.StackOut) == 0 {
		if len(this.StackIn) == 0 {
			return -1
		}
		for len(this.StackIn) > 0 {
			v := this.StackIn[len(this.StackIn)-1]
			this.StackIn = this.StackIn[:len(this.StackIn)-1]
			this.StackOut = append(this.StackOut, v)
		}

	}
	v := this.StackOut[len(this.StackOut)-1]
	this.StackOut = this.StackOut[:len(this.StackOut)-1]
	return v

}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
