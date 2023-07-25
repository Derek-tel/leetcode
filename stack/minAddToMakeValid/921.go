package main

import (
	"fmt"
)

func minAddToMakeValid(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := []rune{}
	for _, v := range s {
		if v == '(' {
			stack = append(stack, v)
		} else if v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack)
}

func test(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := []rune{}

	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack)
}

func get(s string) int {
	var stack []rune

	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack)
}

func four(s string) int {
	var stack []rune
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack)
}

func five(s string) int {
	var stack []rune
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack)
}

func six(s string) int {
	var stack []rune
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack)
}

func seven(s string) int {
	var stack []rune
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack)
}

func eight(s string) int {
	var stack []rune
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack)
}
func main() {
	test := "(()"
	fmt.Println(minAddToMakeValid(test))
}
