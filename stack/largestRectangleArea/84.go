package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	stack := []int{}
	length := len(heights) + 2
	area := 0
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return 0
		}
		return heights[i-1]
	}

	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			weigh := i - stack[len(stack)-1] - 1
			area = max(area, weigh*getHeight(top))
		}
		stack = append(stack, i)
	}
	return area
}

func get(height []int) int {
	var stack []int
	length := len(height) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return height[i-1]
	}
	area := 0
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area = max(area, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return area
}

func three(height []int) int {
	var stack []int
	length := len(height) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		} else {
			return height[i-1]
		}
	}
	maxInt := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	result := 0
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = maxInt(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func four(height []int) int {
	var stack []int
	length := len(height) + 2

	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return height[i-1]
	}
	maxInt := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	result := 0
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = maxInt(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func five(heights []int) int {
	var stack []int
	length := len(heights) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return heights[i-1]
	}
	maxInt := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	result := 0
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = maxInt(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func six(heights []int) int {
	var stack []int
	length := len(heights) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return heights[i-1]
	}
	result := 0
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = max(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func seven(heights []int) int {
	var stack []int
	length := len(heights) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return heights[i-1]
	}
	result := 0
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = max(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func eight(heights []int) int {
	var stack []int
	length := len(heights) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return heights[i-1]
	}
	result := 0
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = max(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func nine(heights []int) int {
	var stack []int
	length := len(heights) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return heights[i-1]
	}
	result := 0
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = max(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func ten(height []int) int {
	var stack []int
	length := len(height) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return height[i-1]
	}
	var result int
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = max(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func eleven(height []int) int {
	var stack []int
	length := len(height) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return height[i-1]
	}
	var result int
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = max(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func twelve(height []int) int {
	var stack []int
	length := len(height) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return height[i-1]
	}
	var result int
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = max(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func thirteen(height []int) int {
	var stack []int
	length := len(height) + 2
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return height[i-1]
	}
	var result int
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = max(result, getHeight(top)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	st := make([]int, 0, 7/2)
	fmt.Println(cap(st))

	test := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(test))

}
