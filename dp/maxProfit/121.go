package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	stack := []int{}
	maxPrice := 0
	for _, price := range prices {
		if len(stack) == 0 || price < stack[len(stack)-1] {
			stack = append(stack, price)
		} else {
			maxPrice = max(maxPrice, price-stack[len(stack)-1])
		}

	}
	return maxPrice
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, maxPrice := prices[0], 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxPrice {
			maxPrice = prices[i] - min
		}

		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxPrice
}

func test(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > res {
			res = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return res
}

func demo(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > res {
			res = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return res
}

func get(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i]-min > res {
			res = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return res
}
