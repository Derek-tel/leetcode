package MinStack

import (
	"math"
)

type MinStack struct {
	Stack []int
	Min   []int
}

func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min:   []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	top := this.Min[len(this.Min)-1]
	this.Min = append(this.Min, min(top, val))
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.Min = this.Min[:len(this.Min)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

func min(x, j int) int {
	if x < j {
		return x
	}
	return j
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
