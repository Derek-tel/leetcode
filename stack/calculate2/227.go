package main

import (
	"fmt"
)

func calculate(s string) int {
	stack := []int{0}
	pre := '+'
	num := 0
	res := 0
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch pre {
			case '+':
				stack = append(stack, num)
				i++
			case '-':
				stack = append(stack, -num)
				i++
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			pre = ch
			num = 0
		}
	}

	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}

func calculate3(s string) int {
	stack := []int{0}
	pre := '+'
	num := 0
	res := 0

	for i, ch := range s {
		isNumber := ch > '0' && ch < '9'
		if isNumber {
			num = num*10 + int(ch-'0')
		}
		if !isNumber && ch != ' ' || i == len(s)-1 {
			switch pre {
			case '+':
				stack = append(stack, num)
			case '_':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			default:
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}
			num = 0
			pre = ch
		}
	}
	for _, item := range stack {
		res = res + item
	}
	return res
}

func test(s string) int {
	pre := '+'
	stack := []int{0}
	num := 0
	for i, ch := range s {
		isNumber := ch >= '0' && ch <= '9'
		if isNumber {
			num = num*10 + int(ch-'0')
		}
		if !isNumber && ch != ' ' || i == len(s)-1 {
			switch pre {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num

			}
			num = 0
			pre = ch
		}
	}
	res := 0
	for _, i := range stack {
		res += i
	}
	return res
}

func get(s string) int {
	pre := '+'
	stack := []int{0}
	num := 0
	for i, ch := range s {
		isNumber := ch >= '0' && ch <= '9'
		if isNumber {
			num = num*10 + int(ch-'0')
		}
		if !isNumber && ch != ' ' || i == len(s)-1 {
			switch pre {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num

			}
			pre = ch
			num = 0
		}
	}
	res := 0
	for _, i := range stack {
		res += i
	}
	return res
}

func five(s string) int {
	pre := '+'
	stack := []int{0}
	num := 0
	for i, ch := range s {
		isNumber := ch >= '0' && ch <= '9'
		if isNumber {
			num = num*10 + int(ch-'0')
		}
		if !isNumber && ch != ' ' || i == len(s)-1 {
			switch pre {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			pre = ch
			num = 0
		}
	}
	result := 0
	for _, i := range stack {
		result += i
	}
	return result
}

func six(s string) int {
	pre := '+'
	stack := []int{0}
	num := 0
	for i, ch := range s {
		isNumber := ch >= '0' && ch <= '9'
		if isNumber {
			num = num*10 + int(ch-'0')
		}
		if !isNumber && ch != ' ' || i == len(s)-1 {
			switch pre {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			num = 0
			pre = ch
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}

func seven(s string) int {
	pre := '+'
	stack := []int{0}
	num := 0
	for i, ch := range s {
		isNumber := ch >= '0' && ch <= '9'
		if isNumber {
			num = num*10 + int(ch-'0')
		}
		if !isNumber && ch != ' ' || i == len(s)-1 {
			switch pre {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}

			pre = ch
			num = 0
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}

func eight(s string) int {
	pre := '+'
	stack := []int{0}
	num := 0
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch pre {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			pre = ch
			num = 0
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}

func main() {
	test := "3*2*2"
	fmt.Println(calculate(test))
}
