package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	res := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}
	fmt.Println(dp)
	return res
}

func demo(nums []int) int {
	dp := make([]int, len(nums))
	res := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func lis(nums []int) int {
	dp := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func get(nums []int) int {
	dp := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func help(nums []int) int {
	dp := make([]int, len(nums))
	res := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func six(nums []int) int {
	dp := make([]int, len(nums))
	result := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func seven(nums []int) int {
	dp := make([]int, len(nums))
	result := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(dp[i], result)
	}
	return result
}

func eight(nums []int) int {
	dp := make([]int, len(nums))
	result := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(dp[i], result)
	}
	return result
}

func nine(nums []int) int {
	dp := make([]int, len(nums))
	result := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func ten(nums []int) int {
	dp := make([]int, len(nums))
	result := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func eleven(nums []int) int {
	dp := make([]int, len(nums))
	result := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
