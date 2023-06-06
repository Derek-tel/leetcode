package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return -1
	}

	left, right := 1, length-1
	ans := -1
	for left <= right {
		mid := left + (right-left)>>1
		count := 0

		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			ans = mid
		}

	}
	return ans
}

func test(num []int) int {
	length := len(num)
	if length <= 1 {
		return -1
	}
	left, right := 1, length-1
	ans := -1
	for left <= right {
		mid := left + (right-left)>>1
		count := 0

		for _, n := range num {
			if n <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			ans = mid
		}
	}
	return ans
}

func three(nums []int) int {
	length := len(nums)
	if length < 1 {
		return -1
	}
	left, right := 1, length-1
	ans := -1
	for left <= right {
		mid := left + (right-left)>>1
		count := 0

		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			ans = mid
		}
	}
	return ans
}

func four(nums []int) int {
	length := len(nums)
	if length < 1 {
		return -1
	}
	left, right := 1, length-1
	ans := -1
	for left <= right {
		mid := left + (right-left)>>1
		count := 0

		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			ans = mid
		}
	}
	return ans
}

func five(nums []int) int {
	length := len(nums)
	if length < 1 {
		return -1
	}
	left, right := 1, length-1
	ans := -1
	for left <= right {
		mid := left + (right-left)>>1
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			ans = mid
		}
	}
	return ans
}

func main() {
	fmt.Println(findDuplicate([]int{1, 1}))
}
