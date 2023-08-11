package main

import (
	"fmt"
	"math"
)

func maxDemo(nums []int) int {
	var result = math.MinInt

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		result = max(result, dp[i])
	}
	return result
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	nums := []int{1, -1, 2, 3, -4, 0}
	fmt.Println(maxDemo(nums))
}
