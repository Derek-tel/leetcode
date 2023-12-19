package main

func maxProfit309(prices []int) int {
	dp := make([][3]int, len(prices))
	//第一天 持有股票  说明今天或者之前买了 结束后收益
	dp[0][0] = -prices[0]
	//第一天结束 处于冷冻期，今天卖了股票
	dp[0][1] = 0
	//第一天结束 冷冻期结束，空仓 昨天卖了 或者之前就卖了
	dp[0][2] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func demo309(prices []int) int {
	dp := make([][3]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func test1(price []int) int {
	dp := make([][3]int, len(price))
	//持有股票  今天买了 或者之前买了
	dp[0][0] = -price[0]
	//第一天  处理冷冻期，今天把股票买了
	dp[0][1] = 0
	//第一天 没有持有  今天卖了或者 之前卖了
	dp[0][2] = 0

	for i := 1; i < len(price); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-price[i])
		dp[i][1] = dp[i-1][0] + price[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}

	return max(dp[len(price)-1][1], dp[len(price)-1][2])

}

func getPrice(price []int) int {
	dp := make([][3]int, len(price))
	//0 持有股票 今天结束 持有股票  今天买了，或者之前买了
	dp[0][0] = -price[0]
	//1 冻结中  今天结束股票冻结，今天卖了
	dp[0][1] = 0
	//2 没有持有股票   第一天没有持有 今天卖了 或者之前卖了
	dp[0][2] = 0

	for i := 1; i < len(price); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-price[i])
		dp[i][1] = dp[i-1][0] + price[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[len(price)-1][1], dp[len(price)-1][2])
}

func five(price []int) int {
	dp := make([][3]int, len(price))
	//0 持有股票
	dp[0][0] = -price[0]
	//1 冻结中
	dp[0][1] = 0
	//3 没有持有股票
	dp[0][2] = 0

	for i := 1; i < len(price); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-price[i])
		dp[i][1] = dp[i-1][0] + price[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[len(price)-1][1], dp[len(price)-1][2])
}

func six309(price []int) int {
	dp := make([][3]int, len(price))
	//0 持有
	dp[0][0] = -price[0]
	//1 冻结，刚卖了
	dp[0][1] = 0
	//2 不持有
	dp[0][2] = 0

	for i := 1; i < len(price); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-price[i])
		dp[i][1] = dp[i-1][0] + price[i]
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}
	return max(dp[len(price)-1][2], dp[len(price)-1][1])
}

func seven309(price []int) int {
	dp := make([][3]int, len(price))
	//0 持有
	dp[0][0] = -price[0]
	//1 冻结，刚卖了
	dp[0][1] = 0
	//2 不持有
	dp[0][2] = 0
	for i := 1; i < len(price); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-price[i])
		dp[i][1] = dp[i-1][0] + price[i]
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}
	return max(dp[len(price)-1][2], dp[len(price)-1][1])
}

func eight309(price []int) int {
	dp := make([][3]int, len(price))
	//0 持有
	dp[0][0] = -price[0]
	//1 冻结
	dp[0][1] = 0
	//2 不持有
	dp[0][2] = 0

	for i := 1; i < len(price); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-price[i])
		dp[i][1] = dp[i-1][0] + price[i]
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}
	return max(dp[len(price)-1][2], dp[len(price)-1][1])
}

func nine309(price []int) int {
	dp := make([][3]int, len(price))
	// 0 持有
	dp[0][0] = -price[0]
	// 1 冻结
	dp[0][1] = 0
	// 2 空仓
	dp[0][2] = 0

	for i := 1; i < len(price); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-price[i])
		dp[i][1] = dp[i-1][0] + price[i]
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}
	return max(dp[len(price)-1][2], dp[len(price)-1][1])
}

func ten309(prices []int) int {
	dp := make([][3]int, len(prices))
	//0 持有
	dp[0][0] = -prices[0]
	//1 冻结
	dp[0][1] = 0
	//2 空仓
	dp[0][2] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(max(dp[len(prices)-1][0], dp[len(prices)-1][1]), dp[len(prices)-1][2])
}

func twelve309(prices []int) int {
	dp := make([][3]int, len(prices))
	//0 持有
	dp[0][0] = -prices[0]
	//1 冻结
	dp[0][1] = 0
	//2 空仓
	dp[0][2] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(max(dp[len(prices)-1][0], dp[len(prices)-1][1]), dp[len(prices)-1][2])
}

func thirteen309(price []int) int {
	dp := make([][3]int, len(price))
	//0 持有
	dp[0][0] = -price[0]
	//1 冻结
	dp[0][1] = 0
	//2 空仓
	dp[0][2] = 0

	for i := 1; i < len(price); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-price[i])
		dp[i][1] = dp[i-1][0] + price[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(max(dp[len(price)-1][0], dp[len(price)-1][1]), dp[len(price)-1][2])
}
