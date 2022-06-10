package main

import (
	"fmt"
	"sort"
)

func demo(nums []int) [][]int {
	var result [][]int
	length := len(nums)
	if length == 0 {
		return result
	}
	sort.Ints(nums)
	for i := 1; i < length-1; i++ {
		start := 0
		end := length - 1
		if i > 1 && nums[i] == nums[i-1] {
			start = i - 1
			end = length - 1
		}
		for start < i && end > i {
			sum := nums[i] + nums[start] + nums[end]
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if sum == 0 {
				result = append(result, []int{nums[start], nums[i], nums[end]})
				start++
				end--
			} else if sum < 0 {
				start++
			} else {
				end--
			}
		}
	}
	return result
}

func main() {
	test := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(demo(test))
	test = []int{0}
	fmt.Println(demo(test))
	test = []int{}
	fmt.Println(demo(test))
}
