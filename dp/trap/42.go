package main

import "fmt"

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
	fmt.Println(left)
	fmt.Println(right)
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
