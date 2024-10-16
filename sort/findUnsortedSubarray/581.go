package main

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}
	left := 0
	right := 0
	flag := nums[left]
	for i := 1; i < length; i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[length-1]
	for i := length - 2; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	res := right - left + 1
	if right == 0 {
		return 0
	}
	return res
}

func test(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	left := 0
	right := 0
	flag := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > flag {
			flag = nums[i]
		} else {
			right = i
		}
	}

	flag = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	res := right - left + 1

	if right == 0 {
		return 0
	}
	return res
}

func get(nums []int) int {
	left := 0
	right := 0
	flag := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	res := right - left + 1
	if right == 0 {
		return 0
	}
	return res
}

func four(nums []int) int {
	left := 0
	right := 0
	flag := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	result := right - left + 1
	if right == 0 {
		return 0
	}
	return result
}

func five(nums []int) int {
	left := 0
	right := 0
	flag := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	if right == 0 {
		return 0
	}
	return right - left + 1
}

func six(nums []int) int {
	left, right := 0, 0
	flag := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	if right == 0 {
		return 0
	}
	return right - left + 1
}

func seven(nums []int) int {
	left, right := 0, 0
	flag := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	if right == 0 {
		return 0
	}
	return right - left + 1
}

func eight(nums []int) int {
	left, right := 0, 0
	flag := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}

	flag = nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}

	if right == 0 {
		return 0
	}
	return right - left + 1
}

func nine(num []int) int {
	left, right := 0, 0
	flag := num[0]
	for i := 0; i < len(num); i++ {
		if num[i] >= flag {
			flag = num[i]
		} else {
			right = i
		}
	}
	flag = num[len(num)-1]
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] <= flag {
			flag = num[i]
		} else {
			left = i
		}
	}
	if right == 0 {
		return 0
	}
	return right - left + 1
}

func ten(num []int) int {
	left, right := 0, 0
	flag := num[0]
	for i := 0; i < len(num); i++ {
		if num[i] >= flag {
			flag = num[i]
		} else {
			right = i
		}
	}
	flag = num[len(num)-1]
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] <= flag {
			flag = num[i]
		} else {
			left = i
		}
	}
	if right == 0 {
		return 0
	}
	return right - left + 1
}

func eleven(nums []int) int {
	left, right := 0, 0
	flag := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	if right == 0 {
		return 0
	}
	return right - left + 1
}

func twelve(nums []int) int {
	left, right := 0, 0
	flag := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	if right == 0 {
		return 0
	}
	return right - left + 1
}

func thirteen(nums []int) int {
	left, right := 0, 0
	flag := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	if right == 0 {
		return 0
	}
	return right - left + 1
}

func fourteen(nums []int) int {
	left, right := 0, 0
	flag := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	if right == 0 {
		return 0
	}
	return right - left + 1
}

func fifteen(nums []int) int {
	left, right := 0, 0
	flag := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= flag {
			flag = nums[i]
		} else {
			right = i
		}
	}
	flag = nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= flag {
			flag = nums[i]
		} else {
			left = i
		}
	}
	if right == 0 {
		return 0
	}
	return right - left + 1
}

func main() {
	test := []int{2, 6, 4, 11, 10, 9, 15}
	fmt.Println(six(test))
}
