package main

import "fmt"

func trap(height []int) int {
	left, right, leftMax, rightMax, sum := 0, len(height)-1, 0, 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				sum = sum + (leftMax - height[left])
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				sum = sum + (rightMax - height[right])
			}
			right--
		}
	}
	return sum
}

func test(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	sum := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				sum = sum + (leftMax - height[left])
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				sum = sum + (rightMax - height[right])
			}
			right--
		}
	}
	return sum
}

func get(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	sum := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				sum = sum + (leftMax - height[left])
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				sum = sum + (rightMax - height[right])
			}
			right--
		}
	}
	return sum
}

func four(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	sum := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				sum = sum + (leftMax - height[left])
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				sum = sum + (rightMax - height[right])
			}
			right--
		}
	}
	return sum
}

func five(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	sum := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				sum = sum + (leftMax - height[left])
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				sum = sum + (rightMax - height[right])
			}
			right--
		}
	}
	return sum
}

func main() {
	test := []int{0, 1, 0, 2, 1, 0}
	result := trap(test)
	fmt.Printf("resulut: %v", result)
}
