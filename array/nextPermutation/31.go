package main

import (
	"fmt"
)

func nextPermutation(nums []int) []int {
	length := len(nums)
	i := 0
	j := 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)
	return nums
}

func demo(nums []int) []int {
	length := len(nums)
	i, j := 0, 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i-1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j < i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)
	return nums
}

func test(nums []int) []int {
	length := len(nums)
	i := 0
	j := 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)
	return nums
}

func four(nums []int) []int {
	length := len(nums)
	i, j := 0, 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)
	return nums
}

func five(num []int) []int {
	length := len(num)
	i, j := 0, 0
	for i = length - 2; i >= 0; i-- {
		if num[i] < num[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if num[j] > num[i] {
				break
			}
		}
		swap(num, i, j)
	}
	reverse(num, i+1, length-1)
	return num
}

func six(nums []int) []int {
	length := len(nums)
	i, j := 0, 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if nums[i] < nums[j] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)
	return nums
}

func seven(nums []int) []int {
	length := len(nums)
	i, j := 0, 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if nums[i] > nums[j] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)
	return nums
}

func eight(nums []int) []int {
	length := len(nums)
	i, j := 0, 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if nums[i] < nums[j] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)

	return nums
}

func nine(nums []int) []int {
	length := len(nums)
	i, j := 0, 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if nums[i] < nums[j] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)
	return nums
}

func ten(nums []int) []int {
	length := len(nums)
	i, j := 0, 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if nums[i] < nums[j] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)
	return nums
}

func eleven(nums []int) []int {
	length := len(nums)
	i, j := 0, 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if nums[i] < nums[j] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)
	return nums
}

func twelve(nums []int) {
	length := len(nums)
	i, j := 0, 0
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = length - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, length-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	fmt.Println(nextPermutation([]int{1, 2, 3, 8, 5, 7, 5, 4}))
}
