package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) (result [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < (n-3) && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {

					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {

					}
				} else if sum < target {
					left++
				} else {
					right--
				}

			}
		}
	}
	return
}

func fourSum1(nums []int, target int) (result [][]int) {
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2]+nums[length-3] < target {
			i++
		}
		for j := i + 1; j < length-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > 0 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				j++
			}
			for left, right := j+1, length-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {

					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {

					}
				} else if sum < target {
					left++
				} else {
					right--
				}

			}
		}
	}
	return
}

func get(nums []int, target int) (result [][]int) {
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2]+nums[length-3] < target {
			i++
		}
		for j := i; j < length-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				j++
			}
			for left, right := j+1, length-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {

					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {

					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}

func demo(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 0; i < length-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2]+nums[length-3] < target {
			continue
		}
		for j := i + 1; j < length-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				continue
			}
			for left, right := j+1, length-1; left < right; {
				num := nums[i] + nums[j] + nums[left] + nums[right]
				if num == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
						continue
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
						continue
					}
				} else if num > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}

func derek(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 0; i < length-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2]+nums[length-3] < target {
			continue
		}
		for j := i + 1; j < length-2 && nums[i]+nums[j]+nums[j+2]+nums[j+3] <= target; j++ {
			if j < i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				continue
			}
			for left, right := j+1, length-1; left < right; {
				num := nums[i] + nums[j] + nums[left] + nums[right]
				if num == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
						continue
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
						continue
					}
				} else if num > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}

func sumFour(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 0; i < length-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2]+nums[length-3] < target {
			continue
		}
		for j := i + 1; j < length-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				continue
			}
			for left, right := j+1, length-1; left < right; {
				total := nums[i] + nums[j] + nums[left] + nums[right]
				if total == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
						continue
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
						continue
					}
				} else if total > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}

func six(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 0; i < length-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[length-1]+nums[length-2]+nums[length-3] < target {
			continue
		}
		for j := i + 1; j < length-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				continue
			}
			for left, right := j+1, length-1; left < right; {
				total := nums[i] + nums[j] + nums[left] + nums[right]
				if total == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
						continue
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
						continue
					}
				} else if total < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}

func main() {
	test := []int{-1, 0, 0, 2, 0}
	result := fourSum(test, 2)
	fmt.Printf("resulut: %v", result)

	type t struct {
		UserId int64
	}
	results := make([]t, 0)
	fmt.Println(results)

	fmt.Println(len("ea9ca2a182a0adcccf95906b1e09bf363b0832c99de6455f10961c8141e1d92ba738e97ac411d6170992efced9521a140224cb3a3546d9c0d0049c08280e4f8d85721eca378f480b46ca1b7c62f335790aa5a050736946136aa9e1d47a1a3628"))

	m := make(map[int]int)
	for _, i := range test {
		md, _ := m[i]
		fmt.Println(md)
	}
}
