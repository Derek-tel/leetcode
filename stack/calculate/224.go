package main

func calculate(s string) int {
	result := 0
	length := len(s)
	sign := 1
	stack := []int{sign}
	for i := 0; i < length; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < length && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result = result + sign*num
		}

	}
	return result
}

func calculate1(s string) int {
	reslt := 0
	length := len(s)
	sign := 1
	stack := []int{1}

	for i := 0; i < length; i++ {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < length && s[i] > '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			reslt = reslt + sign*num
		}
	}
	return reslt
}

func test(s string) int {
	length := len(s)

	result := 0
	sign := 1
	stack := []int{sign}

	for i := 0; i < length; i++ {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < length && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result += sign * num
		}
	}
	return result
}

func get(s string) int {
	length := len(s)
	res := 0
	sign := 1
	stack := []int{sign}
	for i := 0; i < length; {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
		default:
			num := 0
			for ; i < length && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			res += sign * num
		}
	}
	return res
}

func five(s string) int {
	result := 0
	sign := 1
	stack := []int{sign}
	for i := 0; i < len(s); {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result = result + sign*num
		}
	}
	return result
}

func six(s string) int {
	result := 0
	sign := 1
	stack := []int{sign}
	for i := 0; i < len(s); {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result = result + sign*num
		}
	}
	return result
}

func seven(s string) int {
	result := 0
	sign := 1
	stack := []int{sign}
	for i := 0; i < len(s); {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i--
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result = result + sign*num
		}
	}
	return result
}

func eight(s string) int {
	result := 0
	sign := 1
	stack := []int{sign}
	for i := 0; i < len(s); {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result = result + sign*num
		}
	}
	return result
}

func nine(s string) int {
	result := 0
	sign := 1
	stack := []int{sign}
	for i := 0; i < len(s); {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result = result + sign*num
		}
	}
	return result
}

func ten(s string) int {
	result := 0
	sign := 1
	stack := []int{sign}
	for i := 0; i < len(s); {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result = result + num*sign
		}
	}
	return result
}

func eleven(s string) int {
	result := 0
	sign := 1
	stack := []int{sign}
	for i := 0; i < len(s); {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result = result + sign*num
		}
	}
	return result
}

func twelve(s string) int {
	result := 0
	sign := 1
	stack := []int{sign}
	for i := 0; i < len(s); {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result = result + num*sign
		}
	}
	return result
}
func main() {
	s := " 2-1 + 2 "
	calculate(s)
}
