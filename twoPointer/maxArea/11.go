package maxArea

func maxArea(height []int) int {
	length := len(height)
	if length <= 1 {
		return 0
	}
	left, right := 0, length-1
	ans := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func get(height []int) int {
	length := len(height)
	if length <= 1 {
		return 0
	}
	left, right := 0, length-1
	ans := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func test(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	left, right := 0, length-1
	ans := 0

	for left < right {
		area := max(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func four(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	left, right := 0, length-1
	ans := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func five(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	left, right := 0, length-1
	resp := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		resp = max(area, resp)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return resp
}

func six(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	left, right := 0, length-1
	result := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		result = max(result, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
