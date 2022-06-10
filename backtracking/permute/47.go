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

func main() {
	test := []int{1, 1, 1, 2}
	fmt.Println(permuteUnique(test))
}
