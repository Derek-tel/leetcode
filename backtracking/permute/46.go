package main

func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	var dfs func([]int, int, []int)
	res := [][]int{}
	dfs = func(nums []int, index int, temp []int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				temp = append(temp, nums[i])
				dfs(nums, index+1, temp)
				temp = temp[:len(temp)-1]
				used[i] = false
			}
		}
	}
	dfs(nums, 0, []int{})
	return res
}

func demo1(nums []int) [][]int {
	used := make([]bool, len(nums))
	var helper func(int)
	var result [][]int
	var temp []int
	helper = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[index] {
				used[index] = true
				temp = append(temp, nums[i])
				helper(index + 1)
				temp = temp[:len(temp)-1]
				used[i] = false
			}
		}
	}
	helper(0)
	return result
}

func test(nums []int) [][]int {
	use := make([]bool, len(nums))
	var dfs func(int)
	temp := []int{}
	res := [][]int{}
	dfs = func(index int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), temp...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if !use[i] {
				use[i] = true
				temp = append(temp, nums[i])
				dfs(index + 1)
				temp = temp[:len(temp)-1]
				use[i] = false
			}
		}
	}
	dfs(0)
	return res
}

func demo(nums []int) [][]int {
	use := make([]bool, len(nums))
	var helper func(int)
	var temp []int
	var result [][]int
	helper = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if !use[i] {
				use[i] = true
				temp = append(temp, nums[i])
				helper(index + 1)
				temp = temp[:len(temp)-1]
				use[i] = false
			}
		}
	}
	helper(0)
	return result
}

func six(nums []int) [][]int {
	use := make([]bool, len(nums))
	var temp []int
	var result [][]int
	var helper func(int)
	helper = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
		}
		for i := 0; i < len(nums); i++ {
			if !use[i] {
				use[i] = true
				temp = append(temp, nums[i])
				helper(index + 1)
				temp = temp[:len(temp)-1]
				use[i] = false
			}
		}
	}
	helper(0)
	return result
}

func seven(nums []int) [][]int {
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
			if !use[i] {
				use[i] = true
				temp = append(temp, nums[i])
				helper(index + 1)
				temp = temp[:len(temp)-1]
				use[i] = false
			}
		}
	}
	helper(0)
	return result
}
