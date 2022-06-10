package MinStack

import (
	"math"
)

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	m := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(m, val))
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

type Stack struct {
	stack    []int
	minStack []int
}

func constructor() Stack {
	return Stack{
		stack:    []int{},
		minStack: []int{math.MaxInt},
	}
}

func (s *Stack) Push(val int) {
	s.stack = append(s.minStack, val)
	m := s.minStack[len(s.minStack)-1]
	s.minStack = append(s.minStack, min(val, m))
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
