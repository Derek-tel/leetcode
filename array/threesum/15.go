package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, add, length := make([][]int, 0), 0, 0, 0, len(nums)
	for index := 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			add = nums[start] + nums[index] + nums[end]
			if add == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if add < 0 {
				start++
			} else {
				end--
			}
		}
	}
	return result
}

func test(nums []int, target int) (result [][]int) {
	length := len(nums)
	sort.Ints(nums)
	for i := 1; i < length-1; i++ {
		start := 0
		end := length - 1
		if nums[i] == nums[i-1] {
			start = i - 1
		}
		for start < i && end > i {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if sum := nums[start] + nums[i] + nums[end]; sum == target {
				result = append(result, []int{nums[start], nums[i], nums[end]})
			} else if sum > target {
				end--
			} else {
				start++
			}
		}
	}
	return
}

func get(nums []int, target int) (result [][]int) {
	sort.Ints(nums)
	length := len(nums)
	for i := 1; i < length-1; i++ {
		start := 0
		end := length - 1
		if nums[i] == nums[i-1] {
			start = i - 1
		}
		for start < i && end > i {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if sum := nums[start] + nums[i] + nums[end]; sum == target {
				result = append(result, []int{nums[start], nums[i], nums[end]})
			} else if sum > target {
				end--
			} else {
				start++
			}
		}
	}
	return
}

func four(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 1; i < length-1; i++ {
		start, end := 0, length-1
		if i > 1 && nums[i] == nums[i-1] {
			start = i - 1
		}
		for start < i && end > i {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if sum := nums[start] + nums[i] + nums[end]; sum == 0 {
				result = append(result, []int{nums[start], nums[i], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

func five(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 1; i < length; i++ {
		start, end := 0, length-1
		if i > 1 && nums[i] == nums[i-1] {
			start = i - 1
		}
		for start < i && end > i {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if sum := nums[start] + nums[i] + nums[end]; sum == 0 {
				result = append(result, []int{nums[start], nums[i], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

func six(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 1; i < length; i++ {
		start, end := 0, length-1
		if i > 1 && nums[i] == nums[i-1] {
			start = i - 1
		}
		for start < i && end > i {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if sum := nums[start] + nums[i] + nums[end]; sum == 0 {
				result = append(result, []int{nums[start], nums[i], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

func seven(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 0; i < length-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2] < 0 {
			continue
		}
		for left, right := i+1, length-1; left < right; {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
					continue
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
					continue
				}
			} else if total > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func eight(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 0; i < length-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2] < 0 {
			continue
		}
		for left, right := i+1, length-1; left < right; {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
					continue
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
					continue
				}
			} else if total > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func nine(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 0; i < length-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2] < 0 {
			continue
		}
		for left, right := i+1, length-1; left < right; {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
					continue
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
					continue
				}
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func ten(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int

	for i := 0; i < length-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2] < 0 {
			continue
		}
		for left, right := i+1, length-1; left < right; {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
					continue
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
					continue
				}
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func eleven(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int

	for i := 0; i < length-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2] < 0 {
			continue
		}
		for left, right := i+1, length-1; left < right; {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
					continue
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
					continue
				}
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func twelve(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int

	for i := 0; i < length-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2] < 0 {
			continue
		}
		for left, right := i+1, length-1; left < right; {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
					continue
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
					continue
				}
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func thirteen(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int

	for i := 0; i < length-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2] < 0 {
			continue
		}
		for left, right := i+1, length-1; left < right; {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
					continue
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
					continue
				}
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func fourteen(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int

	for i := 0; i < length-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2] < 0 {
			continue
		}
		for left, right := i+1, length-1; left < right; {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
					continue
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
					continue
				}
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func fifteen(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int

	for i := 0; i < length-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2] < 0 {
			continue
		}
		for left, right := i+1, length-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
					continue
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
					continue
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	test := []int{0, 0, 0, 0, 0}
	result := threeSum(test)
	fmt.Printf("resulut: %v", result)
}
