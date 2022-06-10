package main

import "fmt"

func subsets(nums []int) [][]int {
	var dfs func(int)
	set := []int{}
	res := [][]int{}
	dfs = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int(nil), set...))
			return
		}

		//选
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		//不选
		dfs(cur + 1)
	}
	dfs(0)
	return res
}

func demo1(nums []int) [][]int {
	var helper func(int)
	set := []int{}
	var result [][]int
	helper = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), set...))
			return
		}

		set = append(set, nums[index])
		helper(index + 1)
		set = set[:len(set)-1]

		helper(index + 1)
	}
	helper(0)
	return result
}

func test1(nums []int) [][]int {
	var dfs func(int)
	res := [][]int{}
	temp := []int{}
	dfs = func(index int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), temp...))
			return
		}
		//不选
		dfs(index + 1)
		//选择
		temp = append(temp, nums[index])
		dfs(index + 1)
		temp = temp[:len(temp)-1]
	}
	dfs(0)
	return res
}

func demo(nums []int) [][]int {
	var helper func(int)
	var result [][]int
	var temp []int
	helper = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		helper(index + 1)
		temp = append(temp, nums[index])
		helper(index + 1)
		temp = temp[:len(temp)-1]
	}
	helper(0)
	return result
}

func main() {
	test := []int{1, 2, 3}
	fmt.Println(test1(test))
}
