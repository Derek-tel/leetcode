package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	res := []int{}

	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	fmt.Println(nums)
	for i, num := range nums {
		if num > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func get(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	var res []int
	for i, num := range nums {
		if num > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func test(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func four(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func five(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func six(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func seven(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func eight(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func nine(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func ten(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
