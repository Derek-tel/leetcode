package main

import (
	"fmt"
)

func trap(height []int) int {
	stack := []int{}
	sum := 0
	for index, h := range height {
		for len(stack) != 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			seek := stack[len(stack)-1]
			curWidth := index - seek - 1
			sum = sum + curWidth*(min(h, height[seek])-height[top])

		}
		stack = append(stack, index)
	}
	return sum
}

func trap2(height []int) int {
	var stack []int
	sum := 0
	for index, h := range height {
		for len(stack) != 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			seek := stack[len(stack)-1]
			curWeight := index - seek - 1
			sum = sum + curWeight*(min(h, height[seek])-height[top])
		}
		stack = append(stack, index)
	}
	return sum
}

func three(height []int) int {
	var stack []int
	result := 0
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			seek := stack[len(stack)-1]
			weight := index - seek - 1
			result = result + (min(high, height[seek])-height[top])*weight
		}
		stack = append(stack, index)
	}
	return result
}

func four(height []int) int {
	var stack []int
	result := 0
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			seek := stack[len(stack)-1]
			weight := index - seek - 1
			result = result + (min(high, height[seek])-height[top])*weight
		}
		stack = append(stack, index)
	}
	return result
}

func five(height []int) int {
	var stack []int
	result := 0
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			seek := stack[len(stack)-1]
			weight := index - seek - 1
			result = result + (min(high, height[seek])-height[top])*weight
		}
		stack = append(stack, index)
	}
	return result
}

func six(height []int) int {
	var stack []int
	result := 0
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			prev := stack[len(stack)-1]
			weight := index - prev - 1
			result = result + (min(high, height[prev])-height[top])*weight
		}
		stack = append(stack, index)
	}
	return result
}

func seven(height []int) int {
	var stack []int
	var result = 0
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			prev := stack[len(stack)-1]
			weight := index - prev - 1
			result = result + (min(high, height[prev])-height[top])*weight
		}
		stack = append(stack, index)
	}
	return result
}

func eight(height []int) int {
	var stack []int
	var result = 0
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			prev := stack[len(stack)-1]
			weight := index - prev - 1
			result = result + (min(high, height[prev])-height[top])*weight
		}
		stack = append(stack, index)
	}
	return result
}

func nine(height []int) int {
	var stack []int
	var result = 0
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			prev := stack[len(stack)-1]
			weight := index - prev - 1
			result = result + (min(high, height[prev])-height[top])*weight
		}
		stack = append(stack, index)
	}
	return result
}

func ten(height []int) int {
	var stack []int
	var result = 0
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			prev := stack[len(stack)-1]
			weight := index - prev - 1
			result = result + (min(high, height[prev])-height[top])*weight
		}
		stack = append(stack, index)
	}
	return result
}

func eleven(height []int) int {
	var stack []int
	var result = 0
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			prev := stack[len(stack)-1]
			weight := index - prev - 1
			result = result + (min(high, height[prev])-height[top])*weight
		}
		stack = append(stack, index)
	}
	return result
}

func twelve(height []int) int {
	var stack []int
	var result int
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			prev := stack[len(stack)-1]
			weight := index - prev - 1
			result += (min(high, height[prev]) - height[top]) * weight
		}
		stack = append(stack, index)
	}
	return result
}

func thirteen(height []int) int {
	var stack []int
	var result int
	for index, high := range height {
		for len(stack) > 0 && high > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			prev := stack[len(stack)-1]
			weight := index - prev - 1
			result += (min(high, height[prev]) - height[top]) * weight
		}
		stack = append(stack, index)
	}
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	test := []int{0, 1, 0, 2, 1, 0}
	result := trap(test)
	fmt.Printf("resulut: %v", result)
}
