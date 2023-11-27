package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	visit := make([]bool, length)
	perm := []int{}
	res := [][]int{}
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == length {
			res = append(res, append([]int(nil), perm...))
			return
		}

		for i := 0; i < length; i++ {
			if visit[i] || i > 0 && visit[i-1] && nums[i] == nums[i-1] {
				continue
			}
			perm = append(perm, nums[i])
			visit[i] = true
			backtrack(idx + 1)
			visit[i] = false
			perm = perm[:len(perm)-1]
		}

	}
	backtrack(0)
	return res
}

func test1(nums []int) [][]int {
	sort.Ints(nums)
	visit := make([]bool, len(nums))
	temp := []int{}
	res := [][]int{}
	var backTrack func(int)

	backTrack = func(index int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), temp...))
		}

		for i := 0; i < len(nums); i++ {
			if visit[i] || i > 0 && visit[i-1] && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			visit[i] = true
			backTrack(index + 1)
			visit[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	backTrack(0)
	return res
}

func three(nums []int) [][]int {
	sort.Ints(nums)
	use := make([]bool, len(nums))
	var temp []int
	var result [][]int
	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
		}
		for i := 0; i < len(nums); i++ {
			if use[i] || i > 0 && use[i-1] && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			use[i] = true
			backtrack(index + 1)
			use[i] = false
			temp = temp[:len(temp)-1]

		}
	}
	backtrack(0)
	return result
}

func four(nums []int) [][]int {
	sort.Ints(nums)
	use := make([]bool, len(nums))
	var temp []int
	var result [][]int
	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if use[i] || i > 0 && use[i] && nums[i-1] == nums[i-1] {
				continue
			}
			use[i] = true
			temp = append(temp, nums[i])
			backtrack(index + 1)
			temp = temp[:len(temp)-1]
			use[i] = false
		}
	}
	backtrack(0)
	return result
}

func five(nums []int) [][]int {
	sort.Ints(nums)
	use := make([]bool, len(nums))
	var temp []int
	var result [][]int
	var helper func(int)
	helper = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if use[i] || i > 0 && use[i-1] && nums[i] == nums[i-1] {
				continue
			}
			use[i] = true
			temp = append(temp, nums[i])
			helper(index + 1)
			temp = temp[:len(temp)-1]
			use[i] = false
		}
	}
	helper(0)
	return result
}

func eve(nums []int) [][]int {
	sort.Ints(nums)
	use := make([]bool, len(nums))
	var temp []int
	var result [][]int
	var helper func(int)
	helper = func(i int) {
		if i == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		for j := 0; j < len(nums); j++ {
			if use[j] || j > 0 && nums[j] == nums[j-1] && use[j-1] {
				continue
			}
			use[j] = true
			temp = append(temp, nums[j])
			helper(i + 1)
			temp = temp[:len(temp)-1]
			use[j] = false
		}
	}
	helper(0)
	return result
}

func twelve(nums []int) [][]int {
	sort.Ints(nums)
	use := make([]bool, len(nums))
	var temp []int
	var result [][]int
	var handler func(int)
	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if use[i] || i > 0 && nums[i] == nums[i-1] && use[i-1] {
				continue
			}
			use[i] = true
			temp = append(temp, nums[i])
			handler(index + 1)
			temp = temp[:len(temp)-1]
			use[i] = false
		}
	}
	handler(0)
	return result
}

func fourteen(nums []int) [][]int {
	sort.Ints(nums)
	use := make([]bool, len(nums))
	var temp []int
	var result [][]int
	var handler func(int)
	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if use[i] || i > 0 && nums[i] == nums[i-1] && use[i-1] {
				continue
			}
			use[i] = true
			temp = append(temp, nums[i])
			handler(index + 1)
			temp = temp[:len(temp)-1]
			use[i] = false
		}
	}
	handler(0)
	return result
}

func sixteen(nums []int) [][]int {
	sort.Ints(nums)
	use := make([]bool, len(nums))
	var temp []int
	var result [][]int
	var handler func(int)
	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if use[i] || i > 0 && nums[i] == nums[i-1] && use[i-1] {
				continue
			}
			use[i] = true
			temp = append(temp, nums[i])
			handler(index + 1)
			temp = temp[:len(temp)-1]
			use[i] = false
		}
	}

	handler(0)
	return result
}

func eighteen(nums []int) [][]int {
	sort.Ints(nums)
	use := make([]bool, len(nums))
	var temp []int
	var result [][]int
	var handler func(int)
	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if use[i] || i > 0 && nums[i] == nums[i-1] && use[i-1] {
				continue
			}
			use[i] = true
			temp = append(temp, nums[i])
			handler(index + 1)
			temp = temp[:len(temp)-1]
			use[i] = false
		}
	}

	handler(0)
	return result
}

func main() {
	test := []int{1, 1, 1, 2}
	fmt.Println(permuteUnique(test))
}
