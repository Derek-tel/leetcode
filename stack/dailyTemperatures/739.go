package dailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	if length == 0 {
		return []int{}
	}

	stack := []int{}
	answer := make([]int, length)
	for i := 0; i < length; i++ {
		temp := temperatures[i]

		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temp {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[prev] = i - prev
		}

		stack = append(stack, i)
	}
	return answer

}

func test(temperatures []int) []int {
	length := len(temperatures)
	if length == 0 {
		return temperatures
	}
	stack := []int{}
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		temp := temperatures[i]
		for len(stack) != 0 && temp > stack[len(stack)-1] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prev] = i - prev
		}

		stack = append(stack, i)
	}

	return ans
}

func get(temperatures []int) []int {
	length := len(temperatures)
	if length == 0 {
		return temperatures
	}
	stack := []int{}
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		today := temperatures[i]
		for len(stack) != 0 && today > stack[len(stack)-1] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return ans
}

func four(temperatures []int) []int {
	if len(temperatures) == 0 {
		return temperatures
	}
	stack := []int{}
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		today := temperatures[i]
		for len(stack) != 0 && today > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return result
}

func five(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {
		today := temperatures[i]
		for len(stack) > 0 && today > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return result
}

func six(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {
		today := temperatures[i]
		for len(stack) > 0 && today > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return result
}

func seven(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {
		today := temperatures[i]
		for len(stack) > 0 && today > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return result
}

func eight(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {
		today := temperatures[i]
		for len(stack) > 0 && today > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return result
}

func nine(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {
		today := temperatures[i]
		for len(stack) > 0 && today > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return result
}
