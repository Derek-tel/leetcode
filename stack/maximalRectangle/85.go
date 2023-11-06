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

func six(matrix [][]byte) int {
	length := len(matrix[0])
	if length == 0 {
		return 0
	}
	height := make([]int, length)
	var handler func([]int) int
	result := 0
	handler = func(h []int) int {
		res := 0
		long := len(h) + 2
		stack := []int{}
		getHeight := func(i int) int {
			if i == 0 || i == long-1 {
				return -1
			}
			return h[i-1]
		}
		for i := 0; i < long; i++ {
			for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				area := (i - stack[len(stack)-1] - 1) * getHeight(top)
				res = max(res, area)
			}
			stack = append(stack, i)
		}
		return res
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < length; j++ {
			if matrix[i][j] == '1' {
				height[j] += 1
			} else {
				height[j] = 0
			}
		}
		result = max(result, handler(height))
	}
	return result
}

func seven(matrix [][]byte) int {
	var handler func([]int) int
	handler = func(h []int) int {
		length := len(h) + 2
		stack := []int{}
		res := 0
		getHeight := func(i int) int {
			if i == 0 || i == length-1 {
				return -1
			}
			return h[i-1]
		}
		for i := 0; i < length; i++ {
			for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				area := getHeight(top) * (i - stack[len(stack)-1] - 1)
				res = max(res, area)
			}
			stack = append(stack, i)
		}
		return res
	}
	high := make([]int, len(matrix[0]))
	var result int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				high[j] = high[j] + 1
			} else {
				high[j] = 0
			}
		}
		result = max(result, handler(high))
	}
	return result
}

func eight(matrix [][]byte) int {
	var handler func([]int) int
	handler = func(h []int) int {
		length := len(h) + 2
		stack := []int{}
		res := 0
		getHeight := func(i int) int {
			if i == 0 || i == length-1 {
				return -1
			}
			return h[i-1]
		}
		for i := 0; i < length; i++ {
			for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				area := getHeight(top) * (i - stack[len(stack)-1] - 1)
				res = max(res, area)
			}
			stack = append(stack, i)
		}
		return res
	}

	high := make([]int, len(matrix[0]))
	var result int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				high[j] = high[j] + 1
			} else {
				high[j] = 0
			}
		}
		result = max(result, handler(high))
	}
	return result
}

func nine(matrix [][]byte) int {
	handler := func(h []int) int {
		length := len(h) + 2
		stack := []int{}
		result := 0
		getHigh := func(i int) int {
			if i == 0 || i == length-1 {
				return -1
			}
			return h[i-1]
		}
		for i := 0; i < length; i++ {
			for len(stack) > 0 && getHigh(i) < getHigh(stack[len(stack)-1]) {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				area := getHigh(top) * (i - stack[len(stack)-1] - 1)
				result = max(result, area)
			}
			stack = append(stack, i)
		}
		return result
	}
	high := make([]int, len(matrix[0]))
	var v int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				high[j] = high[j] + 1
			} else {
				high[j] = 0
			}
		}
		v = max(v, handler(high))
	}
	return v
}

func ten(matrix [][]byte) int {
	handler := func(h []int) int {
		length := len(h) + 2
		stack := []int{}
		result := 0
		getHigh := func(i int) int {
			if i == 0 || i == length-1 {
				return -1
			}
			return h[i-1]
		}
		for i := 0; i < length; i++ {
			for len(stack) > 0 && getHigh(i) < getHigh(stack[len(stack)-1]) {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result = max(result, getHigh(top)*(i-stack[len(stack)-1]-1))
			}
			stack = append(stack, i)
		}
		return result
	}

	high := make([]int, len(matrix[0]))
	var v int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				high[j] += 1
			} else {
				high[j] = 0
			}
		}
		v = max(v, handler(high))
	}
	return v
}

func eleven(matrix [][]byte) int {
	var handler func([]int) int
	handler = func(h []int) int {
		stack := []int{}
		result := 0
		length := len(h) + 2
		getHeight := func(i int) int {
			if i == 0 || i == length-1 {
				return -1
			}
			return h[i-1]
		}
		for i := 0; i < length; i++ {
			for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result = max(result, getHeight(top)*(i-stack[len(stack)-1]-1))
			}
			stack = append(stack, i)
		}
		return result
	}

	var high = make([]int, len(matrix[0]))
	var resp = 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j]-'0' == 1 {
				high[j] += 1
			} else {
				high[j] = 0
			}
		}
		resp = max(resp, handler(high))
	}
	return resp
}

func twelve(matrix [][]byte) int {
	var handler func([]int) int
	handler = func(h []int) int {
		result := 0
		stack := make([]int, 0)
		length := len(h) + 2
		getHeight := func(i int) int {
			if i == 0 || i == length-1 {
				return -1
			}
			return h[i-1]
		}
		for i := 0; i < length; i++ {
			for len(stack) > 0 && getHeight(i) < getHeight(stack[len(stack)-1]) {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result = max(result, getHeight(top)*(i-stack[len(stack)-1]-1))
			}
			stack = append(stack, i)
		}
		return result
	}

	var high = make([]int, len(matrix[0]))
	var resp = 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j]-'0' == 0 {
				high[j] = 0
			} else {
				high[j] += 1
			}
		}
		resp = max(resp, handler(high))
	}
	return resp
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
