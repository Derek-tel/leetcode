package main

import (
	"fmt"
)

func trap(height []int) int {
	highest := 0
	left := make([]int, len(height))
	right := make([]int, len(height))
	sum := 0
	for i := 0; i < len(height); i++ {
		if height[i] > highest {
			highest = height[i]
		} else {
			left[i] = highest - height[i]
		}
	}
	highest = 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > highest {
			highest = height[i]
		} else {
			right[i] = highest - height[i]
		}
	}
	for i := range height {
		sum = sum + min(left[i], right[i])
	}
	return sum
}

func test(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	high := 0

	for i := 0; i < len(height); i++ {
		if height[i] > high {
			high = height[i]
		} else {
			left[i] = height[i]
		}
	}
	high = 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > high {
			high = height[i]
		} else {
			right[i] = height[i]
		}
	}
	sum := 0
	for i, _ := range height {
		sum = sum + min(left[i], right[i])
	}

	return sum

}

func get(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	high := 0
	for i := 0; i < len(height); i++ {
		if height[i] > high {
			high = height[i]
		} else {
			left[i] = height[i]
		}
	}
	high = 0
	for j := len(height) - 1; j >= 0; j-- {
		if height[j] > high {
			high = height[j]
		} else {
			right[j] = height[j]
		}
	}
	sum := 0
	for i := 0; i < len(height); i++ {
		sum += min(left[i], right[i])
	}
	return sum
}

func four(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[len(height)-1] = height[len(height)-1]
	for j := len(height) - 2; j >= 0; j-- {
		right[j] = max(right[j+1], height[j])
	}
	result := 0
	for i := 0; i < len(height); i++ {
		result += min(left[i], right[i]) - height[i]
	}
	return result
}

func five(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	result := 0
	for i := 0; i < len(height); i++ {
		result += min(left[i], right[i]) - height[i]
	}
	return result
}

func six(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	result := 0
	for i := 0; i < len(height); i++ {
		result += min(left[i], right[i]) - height[i]
	}
	return result
}

func seven(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	result := 0
	for i := 0; i < len(height); i++ {
		result += min(left[i], right[i]) - height[i]
	}
	return result
}

func eight(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	result := 0
	for i := 0; i < len(height); i++ {
		result += min(left[i], right[i]) - height[i]
	}
	return result
}

func nine(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}

	right[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	result := 0
	for i := 0; i < len(height); i++ {
		result += min(left[i], right[i]) - height[i]
	}
	return result
}

func ten(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[len(right)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	result := 0
	for i := 0; i < len(height); i++ {
		result += min(left[i], right[i]) - height[i]
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	test := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := trap(test)
	fmt.Printf("resulut: %v", result)
}
