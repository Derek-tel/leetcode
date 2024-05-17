package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	dp := make([]bool, target+1)
	res := make(map[int][][]int)
	dictMap := make(map[int]bool)
	for _, nums := range candidates {
		dictMap[nums] = true
	}
	dp[0] = true
	for i := 1; i <= target; i++ {
		for j := 0; j < len(candidates); j++ {
			if i >= candidates[j] {
				dp[i] = dp[i] || dp[i-candidates[j]]
				if dp[i-candidates[j]] {
					temp := res[i-candidates[j]]
					comI := res[i]
					if len(temp) == 0 {
						tempi := []int{candidates[j]}
						comI = append(comI, tempi)
					} else {
						for _, com := range temp {
							newCom := append(com, candidates[j])
							comI = append(comI, newCom)
						}
					}

					res[i] = comI
				}
			}
		}
	}
	return res[target]
}

func combinationSum1(candidates []int, target int) [][]int {
	if len(candidates) == 0 || target == 0 {
		return [][]int{}
	}
	res := [][]int{}
	dfs(candidates, target, 0, []int{}, &res)
	return res
}

func dfs(candidates []int, target int, index int, temp []int, res *[][]int) {
	if target == 0 {
		newTemp := []int{}
		for _, v := range temp {
			newTemp = append(newTemp, v)
		}
		*res = append(*res, newTemp)
		return
	}
	if target < 0 {
		return
	}

	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}
		temp = append(temp, candidates[i])
		dfs(candidates, target-candidates[i], i, temp, res)
		temp = temp[:len(temp)-1]
	}
}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func combination(nums []int, tar int) (res [][]int) {
	com := []int{}
	var dfs func(int, int)
	dfs = func(target int, index int) {
		if index == len(nums) {
			return
		}
		if target == 0 {
			res = append(res, append([]int(nil), com...))
			return
		}
		//不访问
		dfs(target, index+1)
		//访问
		if target >= nums[index] {
			com = append(com, nums[index])
			dfs(target-nums[index], index)
			com = com[:len(com)-1]
		}
	}
	dfs(tar, 0)
	return
}

func demo(nums []int, target int) [][]int {
	var com []int
	var result [][]int
	var helper func(int, int)
	helper = func(tar int, index int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), com...))
			return
		}
		helper(target, index+1)
		if target > nums[index] {
			com = append(com, nums[index])
			helper(tar-nums[index], index)
			com = com[:len(nums)-1]
		}
	}
	helper(target, 0)
	return result
}

func four(nums []int, target int) [][]int {
	var com []int
	var result [][]int
	var helper func(int, int)
	helper = func(tar int, index int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), com...))
			return
		}
		//不选
		helper(tar, index+1)
		if target > nums[index] {
			com = append(com, nums[index])
			helper(tar-nums[index], index)
			com = com[:len(com)-1]
		}
	}
	helper(target, 0)
	return result
}

func five(nums []int, target int) [][]int {
	var com []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), com...))
			return
		}
		handler(index+1, tar)

		if tar > nums[index] {
			com = append(com, nums[index])
			handler(index, tar-nums[index])
			com = com[:len(com)-1]
		}
	}
	handler(0, target)
	return result
}

func six(nums []int, target int) [][]int {
	var com []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), com...))
			return
		}

		//不选
		handler(index+1, tar)

		//选
		if tar >= nums[index] {
			com = append(com, nums[index])
			handler(index, tar-nums[index])
			com = com[:len(com)-1]
		}
	}
	handler(0, target)
	return result
}

func seven(nums []int, target int) [][]int {
	var com []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), com...))
			return
		}

		handler(index+1, tar)

		if tar >= nums[index] {
			com = append(com, nums[index])
			handler(index, tar-nums[index])
			com = com[:len(com)-1]
		}
	}
	handler(0, target)
	return result
}

func eight(nums []int, target int) [][]int {
	var temp []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), temp...))
			return
		}

		handler(index+1, tar)

		if tar >= nums[index] {
			temp = append(temp, nums[index])
			handler(index, tar-nums[index])
			temp = temp[:len(temp)-1]
		}
	}
	handler(0, target)
	return result
}

func nine(nums []int, target int) [][]int {
	var temp []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), temp...))
			return
		}

		handler(index+1, tar)

		if tar >= nums[index] {
			temp = append(temp, nums[index])
			handler(index, tar-nums[index])
			temp = temp[:len(temp)-1]
		}
	}
	handler(0, target)
	return result
}

func ten(nums []int, target int) [][]int {
	var temp []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), temp...))
			return
		}

		//no
		handler(index+1, tar)

		//yes
		if tar >= nums[index] {
			temp = append(temp, nums[index])
			handler(index, tar-nums[index])
			temp = temp[:len(temp)-1]
		}
	}
	handler(0, target)
	return result
}

func eleven(nums []int, target int) [][]int {
	var temp []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), temp...))
			return
		}

		//no
		handler(index+1, tar)

		//yes
		if tar >= nums[index] {
			temp = append(temp, nums[index])
			handler(index, tar-nums[index])
			temp = temp[:len(temp)-1]
		}
	}

	handler(0, target)
	return result
}

func twelve(nums []int, target int) [][]int {
	var temp []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), temp...))
			return
		}

		handler(index+1, tar)

		if tar >= nums[index] {
			temp = append(temp, nums[index])
			handler(index, tar-nums[index])
			temp = temp[:len(temp)-1]
		}
	}
	handler(0, target)
	return result
}

func thirteen(nums []int, target int) [][]int {
	var temp []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, target int) {
		if index == len(nums) {
			return
		}
		if target == 0 {
			result = append(result, append([]int(nil), temp...))
			return
		}
		handler(index+1, target)

		if target >= nums[index] {
			temp = append(temp, nums[index])
			handler(index, target-nums[index])
			temp = temp[:len(temp)-1]
		}
	}
	handler(0, target)
	return result
}

func fourteen(nums []int, target int) [][]int {
	var temp []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), temp...))
			return
		}
		handler(index+1, tar)

		if tar >= nums[index] {
			temp = append(temp, nums[index])
			handler(index, tar-nums[index])
			temp = temp[:len(temp)-1]
		}
	}
	handler(0, target)
	return result
}

func fifteen(nums []int, target int) [][]int {
	var temp []int
	var result [][]int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if index == len(nums) {
			return
		}
		if tar == 0 {
			result = append(result, append([]int(nil), temp...))
			return
		}

		handler(index+1, tar)

		if nums[index] <= tar {
			temp = append(temp, nums[index])
			handler(index, tar-nums[index])
			temp = temp[:len(temp)-1]
		}
	}
	handler(0, target)
	return result
}

func main() {
	test := []int{4, 3, 2, -2}
	tar := 7
	fmt.Println(fifteen(test, tar))
}
