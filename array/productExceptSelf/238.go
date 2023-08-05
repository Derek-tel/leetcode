package main

import "fmt"

func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	r := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * r
		r = r * nums[i]
	}
	return answer
}

func test(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	r := 1
	for j := length - 1; j >= 0; j-- {
		answer[j] = answer[j] * r
		r = r * nums[j]
	}
	return answer
}

func three(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i] - 1*answer[i-1]
	}
	r := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * r
		r = r * nums[i]
	}
	return answer
}

func four(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	r := 1
	for j := length - 1; j >= 0; j-- {
		result[j] = result[j] * r
		r = r * nums[j]
	}
	return result
}

func five(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	r := 1
	for j := length - 1; j >= 0; j-- {
		result[j] = result[j] * r
		r = r * nums[j]
	}
	return result
}

func six(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	r := 1
	for j := length - 1; j >= 0; j-- {
		result[j] = result[j] * r
		r = r * nums[j]
	}
	return result
}

func seven(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	r := 1
	for j := length - 1; j >= 0; j-- {
		result[j] = result[j] * r
		r = r * nums[j]
	}
	return result
}

func eight(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	r := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= r
		r = r * nums[i]
	}
	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
