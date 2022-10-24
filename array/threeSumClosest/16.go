package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	result := 0
	diff := math.MaxInt
	if length > 2 {
		sort.Ints(nums)
		for i := 0; i < length-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, length-1; j < k; {
				sum := nums[i] + nums[k] + nums[j]
				if abs(sum-target) < diff {
					result, diff = sum, abs(sum-target)
				}
				if sum == target {
					return result
				}
				if sum < target {
					j++
				} else {
					k--
				}
			}
		}
	}
	return result
}

func test(nums []int, target int) (result int) {
	length := len(nums)
	flag := math.MaxInt
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for start, end := i+1, length-1; start < end; {
			sum := nums[i] + nums[start] + nums[end]
			if abs(sum-target) < flag {
				result, flag = sum, abs(sum-target)
			}
			if sum == target {
				return sum
			} else if sum > target {
				end--
			} else {
				start++
			}
		}
	}
	return
}

func get(nuns []int, target int) (result int) {
	length := len(nuns)
	sort.Ints(nuns)
	flag := math.MaxInt
	for i := 0; i < length; i++ {
		if i > 0 && nuns[i] == nuns[i-1] {
			continue
		}
		for start, end := i+1, length-1; start < end; {
			sum := nuns[i] + nuns[start] + nuns[end]
			if abs(sum-target) < flag {
				result, flag = sum, abs(sum-target)
			}
			if sum == target {
				return sum
			} else if sum > target {
				end--
			} else {
				start++
			}
		}
	}
	return
}

func four(nums []int, target int) int {
	length := len(nums)
	var result int
	sort.Ints(nums)
	flag := math.MaxInt
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i]-1 {
			continue
		}
		for start, end := i+1, length-1; start < end; {
			sum := nums[i] + nums[start] + nums[end]
			if abs(sum-target) < flag {
				result, flag = sum, abs(sum-target)
			}
			if sum == target {
				return sum
			} else if sum < target {
				start++
			} else {
				end--
			}
		}
	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	test := []int{-1, 0, 0, 1, 0}
	result := threeSumClosest(test, 2)
	fmt.Printf("resulut: %v", result)
}
