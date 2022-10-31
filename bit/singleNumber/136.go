package main

import "fmt"

func singleNumber(nums []int) int {
	length := len(nums)
	if length < 2 {
		return nums[1]
	}
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}

func test(nums []int) int {
	res := 0
	for _, num := range nums {
		res = num ^ res
	}
	return res
}

func three(nums []int) int {
	res := 0
	for _, num := range nums {
		res = num ^ res
	}
	return res
}

func four(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
