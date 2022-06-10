package main

import "fmt"

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left, right := 0, length-1

	for left <= right {
		half := (left + right) / 2
		if nums[half] == target {
			return half
		}
		if nums[0] <= nums[half] {
			if nums[0] <= target && target < nums[half] {
				right = half - 1
			} else {
				left = half + 1
			}
		} else {
			if nums[half] < target && nums[length-1] >= target {
				left = half + 1
			} else {
				right = half - 1
			}
		}
	}
	return -1
}

func test(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		half := (right + left) / 2
		if nums[half] == target {
			return half
		}
		if nums[0] <= nums[half] {
			if target >= nums[0] && nums[half] > target {
				right = half - 1
			} else {
				left = half + 1
			}
		} else {
			if nums[half] < target && target <= nums[right] {
				left = half + 1
			} else {
				right = half - 1
			}
		}
	}
	return -1
}

func get(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		half := (left + right) / 2
		if nums[half] == target {
			return half
		}
		if nums[0] <= nums[half] {
			if target >= nums[0] && nums[half] > target {
				right = half - 1
			} else {
				left = half + 1
			}
		} else {
			if nums[half] < target && target <= nums[length-1] {
				left = half + 1
			} else {
				right = half - 1
			}
		}
	}
	return -1
}

func main() {
	test := []int{3, 1}
	fmt.Println(search(test, 1))
}
