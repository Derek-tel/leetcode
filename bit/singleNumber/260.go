package main

import "fmt"

func one(nums []int) []int {
	tag := 0
	for i := 0; i < len(nums); i++ {
		tag ^= nums[i]
	}

	flag := tag & (-tag)

	left, right := 0, 0
	for _, num := range nums {
		if num&flag > 0 {
			left ^= num
		} else {
			right ^= num
		}
	}
	return []int{left, right}
}

func two(nums []int) []int {
	tag := 0
	for i := 0; i < len(nums); i++ {
		tag ^= nums[i]
	}

	flag := tag & (-tag)

	left, right := 0, 0
	for _, num := range nums {
		if num&flag > 0 {
			left ^= num
		} else {
			right ^= num
		}
	}
	return []int{left, right}
}

func fourteen(nums []int) []int {
	tag := 0
	for i := 0; i < len(nums); i++ {
		tag ^= nums[i]
	}
	flag := tag & (-tag)

	left, right := 0, 0
	for _, num := range nums {
		if num&flag > 0 {
			left ^= num
		} else {
			right ^= num
		}
	}
	return []int{left, right}
}

func main() {
	demo := 6
	fmt.Println(demo & (-demo))
}
