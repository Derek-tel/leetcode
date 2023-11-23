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

func five(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	var result int
	flag := math.MaxInt
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for start, end := i+1, length-1; start < end; {
			sum := nums[i] + nums[start] + nums[end]
			if abs(sum-target) < flag {
				flag = abs(sum - target)
				result = sum
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

func six(nums []int, target int) int {
	var result int
	sort.Ints(nums)
	flag := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, len(nums)-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < flag {
				flag = abs(sum - target)
				result = sum
			}
			if sum == target {
				return sum
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func seven(nums []int, target int) int {
	var result int
	sort.Ints(nums)
	flag := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, len(nums)-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < flag {
				flag = abs(sum - target)
				result = sum
			}
			if sum == target {
				return sum
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func eight(nums []int, target int) int {
	var result int
	sort.Ints(nums)
	flag := math.MaxInt
	if len(nums) > 2 {
		for i := 0; i < len(nums)-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for left, right := i+1, len(nums)-1; left < right; {
				sum := nums[i] + nums[left] + nums[right]
				if abs(sum-target) < flag {
					flag = abs(sum - target)
					result = sum
				}
				if sum == target {
					return sum
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}

func nine(nums []int, target int) int {
	var result int
	sort.Ints(nums)
	flag := math.MaxInt
	if len(nums) > 2 {
		for i := 0; i < len(nums)-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for left, right := i+1, len(nums)-1; left < right; {
				sum := nums[i] + nums[left] + nums[right]
				if abs(sum-target) < flag {
					flag = abs(sum - target)
					result = sum
				}
				if sum == target {
					return result
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}

func ten(nums []int, target int) int {
	var result int
	sort.Ints(nums)
	flag := math.MaxInt
	if len(nums) > 2 {
		for i := 0; i < len(nums)-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for left, right := i+1, len(nums)-1; left < right; {
				sum := nums[i] + nums[left] + nums[right]
				if abs(sum-target) < flag {
					flag = abs(sum - target)
					result = sum
				}
				if sum == target {
					return result
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}

func eleven(nums []int, target int) int {
	var result int
	sort.Ints(nums)
	flag := math.MaxInt
	if len(nums) > 2 {
		for i := 0; i < len(nums)-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for left, right := i+1, len(nums)-1; left < right; {
				sum := nums[i] + nums[left] + nums[right]
				if abs(sum-target) < flag {
					flag = abs(sum - target)
					result = sum
				}
				if sum == target {
					return result
				} else if sum < target {
					left++
				} else {
					right--
				}
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
	//test := []int{-1, 0, 0, 1, 0}
	//result := six(test, 2)
	//fmt.Printf("resulut: %v", result)
	testDemo := []int{-3, 0, 1, 1, 1, 2}
	result := six(testDemo, 2)
	fmt.Printf("resulut: %v \n", result)

	left := 1

	for left++; left < 0; left++ {
		continue
	}
	fmt.Println(left)

	testDemo = []int{-3, 0, 1, 1, 1, 1, 1, 2}

	fmt.Println(cap(testDemo))

	tag := len(testDemo)
	for i, i2 := range testDemo {
		if i2 == 2 {
			tag = i
		}
	}
	fmt.Println(tag)
	testDemo = append(testDemo[:tag], testDemo[tag+1:]...)
	fmt.Println(testDemo)
}
