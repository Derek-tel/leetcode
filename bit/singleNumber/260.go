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

func three260(nums []int) []int {
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

func four260(nums []int) []int {
	tag := 0
	for i := 0; i < len(nums); i++ {
		tag ^= nums[i]
	}
	flag := tag & (-tag)
	left, right := 0, 0
	for _, n := range nums {
		if n&flag > 0 {
			left ^= n
		} else {
			right ^= n
		}
	}
	return []int{left, right}
}

func main() {
	demo := 6
	fmt.Println(demo & (-demo))
}
