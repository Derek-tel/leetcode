package main

import (
	"fmt"
)

// dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
// j-nums[i] >=0   j>=num[s]
func canPartition(nums []int) bool {
	count := 0
	for _, v := range nums {
		count = count + v
	}

	if count%2 != 0 {
		return false
	}
	half := count / 2
	dp := make([]bool, half+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		for j := half; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	fmt.Printf("%+v", dp)
	return dp[half]
}
func test(nums []int) bool {
	count := 0
	for _, v := range nums {
		count = count + v
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	dp := make([]bool, half+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		for j := half; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[half]
}

func get(nums []int) bool {
	count := 0
	for _, num := range nums {
		count += num
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for j := half; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[half]
}

func demo(nums []int) bool {
	count := 0
	for _, num := range nums {
		count += num
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for j := half; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[half]
}

func five(nums []int) bool {
	count := 0
	max := 0
	for _, num := range nums {
		count += num
		if num > max {
			max = num
		}
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	if max > half {
		return false
	}
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for j := half; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[half]
}

func six(nums []int) bool {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	if max > half {
		return false
	}
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for j := half; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[half]
}

func seven(nums []int) bool {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		count = count + nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	if half < max {
		return false
	}
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for j := half; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[half]
}

func eight(nums []int) bool {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		count = count + nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	if half < max {
		return false
	}
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for j := half; j >= num; j++ {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[half]
}

func nine(nums []int) bool {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	if half < max {
		return false
	}
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for j := half; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[half]
}

func ten(nums []int) bool {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	if half < max {
		return false
	}
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for i := half; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[half]
}

func eleven(nums []int) bool {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	if half < max {
		return false
	}
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for i := half; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[half]
}

func twelve(nums []int) bool {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	if half < max {
		return false
	}

	dp := make([]bool, half+1)
	dp[0] = true

	for _, num := range nums {
		for i := half; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[half]
}

func thirteen(nums []int) bool {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	if half < max {
		return false
	}
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for i := half; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[half]
}

func fourteen(nums []int) bool {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if count%2 == 1 {
		return false
	}
	half := count / 2
	if max > half {
		return false
	}
	dp := make([]bool, half+1)
	dp[0] = true
	for _, num := range nums {
		for i := half; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[half]
}

func main() {
	fmt.Println(canPartition([]int{3, 3, 3, 4, 5}))
}
