package main

func findTargetSumWays(nums []int, target int) int {
	count := 0
	for _, v := range nums {
		count = count + v
	}
	diff := count - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	length, neg := len(nums), diff/2
	dp := make([][]int, length+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i := 0; i < length; i++ {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j > nums[i] {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}

	return dp[length][neg]
}

func findTargetSumWays1(nums []int, target int) int {
	count := 0
	for _, v := range nums {
		count = count + v
	}
	diff := count - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	_, neg := len(nums), diff/2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[neg]
}

func demo(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		count += num
	}
	diff := count - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}

func test(nums []int, target int) int {
	count := 0
	for _, n := range nums {
		count += n
	}
	diff := count - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
		}
	}

	return dp[neg]

}

func get(nums []int, target int) int {
	count := 0
	for _, n := range nums {
		count += n
	}
	diff := count - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for i := neg; i > num; i-- {
			dp[i] = dp[i] + dp[i-num]
		}
	}
	return dp[neg]
}

func six(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		count += num
	}
	diff := count - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for i := neg; i >= num; i-- {
			dp[i] = dp[i] + dp[i-num]
		}
	}
	return dp[neg]
}

func seven(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		count += num
	}
	diff := count - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for i := neg; i >= num; i-- {
			dp[i] = dp[i] + dp[i-num]
		}
	}
	return dp[neg]
}

func eight(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		count += num
	}
	diff := count - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}

	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
		}
	}
	return dp[neg]
}

func nine(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		count += num
	}
	diff := count - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
		}
	}
	return dp[neg]
}
