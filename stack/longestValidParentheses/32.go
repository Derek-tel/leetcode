package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func test(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := []int{-1}
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1]-1)
			}
		}
	}
	return res
}

func get(s string) int {
	stack := []int{-1}
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func four(s string) int {
	stack := []int{-1}
	result := 0
	maxInt := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = maxInt(result, i-stack[len(stack)-1])
			}
		}
	}
	return result
}

func five(s string) int {
	stack := []int{-1}
	result := 0
	maxInt := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			fmt.Println(stack)
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = maxInt(result, i-stack[len(stack)-1])
			}
		}
	}
	return result
}

func six(s string) int {
	stack := []int{-1}
	result := 0
	maxInt := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = maxInt(result, i-stack[len(stack)-1])
			}
		}
	}
	return result
}

func seven(s string) int {
	stack := []int{-1}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = max(result, i-stack[len(stack)-1])
			}
		}
	}
	return result
}

func eight(s string) int {
	stack := []int{-1}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = max(result, i-stack[len(stack)-1])
			}
		}
	}
	return result
}

func nine(s string) int {
	stack := []int{-1}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = max(result, i-stack[len(stack)-1])
			}
		}
	}
	return result
}

func ten(s string) int {
	stack := []int{-1}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = max(result, i-stack[len(stack)-1])
			}
		}
	}
	return result
}

func eleven(s string) int {
	stack := []int{-1}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = max(result, i-stack[len(stack)-1])
			}
		}
	}
	return result
}

func twelve(s string) int {
	stack := []int{-1}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = max(result, i-stack[len(stack)-1])
			}
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(five("(())"))
}
