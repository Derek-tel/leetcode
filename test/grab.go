package main

import (
	"fmt"
	"math"
)

func revert(x int) int {
	if math.MinInt32 > x || x > math.MaxInt32 {
		return 0
	}

	newX := 0
	for x != 0 {
		mod := x % 10
		x = x / 10
		newX = newX*10 + mod
	}
	fmt.Println(newX)
	return newX
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	maxInt := Max(dp[0], dp[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = Max(nums[i]+dp[i-2], dp[i-1])
		maxInt = Max(maxInt, dp[i])
	}
	return maxInt
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
