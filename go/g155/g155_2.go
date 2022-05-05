package main

type MinStack struct {
	Elements []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{Elements: make([]int, 0)}
	// stack.Elements = new([]int)
	return stack
}

func (this *MinStack) Len() int {
	return len(this.Elements)
}

func (this *MinStack) Push(x int) {
	this.Elements = append(this.Elements, x)
	//	fmt.Println(this.Len())
}

func (this *MinStack) Pop() {
	if this.Len() != 0 {
		this.Elements = this.Elements[:this.Len()-1]
	}
}

func (this *MinStack) Top() int {
	if this.Len() != 0 {
		return this.Elements[this.Len()-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.Len() == 0 {
		return 0
	}
	min := this.Elements[0]
	for _, i := range this.Elements {
		if i < min {
			min = i
		}
	}
	return min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
