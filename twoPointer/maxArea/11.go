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

func seven(height []int) int {
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

func eight(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	left, right := 0, length-1
	result := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		result = max(area, result)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func nine(height []int) int {
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

func ten(height []int) int {
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

func eleven(height []int) int {
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

func twelve(height []int) int {
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

func thirteen(height []int) int {
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

func fourteen(height []int) int {
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
