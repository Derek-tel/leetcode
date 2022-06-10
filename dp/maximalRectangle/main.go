package main

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	height := make([]int, len(matrix[0]))
	res := 0
	for row := 0; row < len(matrix); row++ {

		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == '1' {
				height[col] += 1
			} else {
				height[col] = 0
			}
		}
		res = max(res, largestRectangleArea(height))
	}
	return 0
}

func largestRectangleArea(heights []int) int {
	stack := []int{}
	length := len(heights) + 2
	area := 0
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return 0
		}
		return heights[i-1]
	}

	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			weigh := i - stack[len(stack)-1] - 1
			area = max(area, weigh*getHeight(top))
		}
		stack = append(stack, i)
	}
	return area
}

func demo(heights []int) int {
	stack := []int{}
	length := len(heights) + 2
	area := 0
	getHeight := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return heights[i-1]
	}
	for i := 0; i < length; i++ {
		for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			weight := i - stack[len(stack)-1] - 1
			area = max(area, weight*getHeight(top))
		}
		stack = append(stack, i)
	}
	return area
}

func test(matrix [][]byte) int {
	length := len(matrix)
	if length == 0 {
		return 0
	}
	height := make([]int, length)
	res := 0

	for i := 0; i < length; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				height[j] += 1
			} else {
				height[j] = 0
			}
		}
		res = max(res, helper(height))
	}
	return res
}

func helper(height []int) int {
	res := 0
	length := len(height) + 2
	stack := []int{}
	getHigh := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return height[i-1]
	}
	for i := 0; i < length; i++ {
		high := getHigh(i)
		for len(stack) > 0 && high < stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area := (i - stack[len(stack)-1] - 1) * getHigh(top)
			res = max(res, area)
		}
		stack = append(stack, i)
	}
	return res
}

func get(height []int) int {
	res := 0
	length := len(height) + 2
	stack := []int{}
	getHigh := func(i int) int {
		if i == 0 || i == length-1 {
			return -1
		}
		return height[i-1]
	}

	for i := 0; i < length; i++ {
		high := getHigh(i)
		for len(stack) > 0 && high < stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area := (i - stack[len(stack)-1] - 1) * getHigh(top)
			res = max(res, area)
		}

		stack = append(stack, i)
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
