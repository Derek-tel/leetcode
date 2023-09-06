package main

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	generate(n, n, "", &res)
	return res
}

func generate(left int, right int, str string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
		return
	}
	if left > right {
		return
	}
	if left > 0 {
		generate(left-1, right, str+"(", res)
	}
	if right > 0 {
		generate(left, right-1, str+")", res)
	}
}

func demo(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	helper(n, n, "", &res)
	return res
}

func helper(left, right int, str string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
		return
	}
	if left > right {
		return
	}
	if left > 0 {
		helper(left-1, right, str+"(", res)
	}
	if right > 0 {
		helper(left, right-1, str+")", res)
	}
}

func test(n int) []string {
	res := []string{}
	generate1(n, n, "", &res)
	return res
}

func generate1(left int, right int, str string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
		return
	}
	if left > right {
		return
	}
	if left > 0 {
		generate(left-1, right, str+"(", res)
	}
	if right > 0 {
		generate(left, right-1, str+")", res)
	}
}

func four(n int) []string {
	var result []string
	var generator func(int, int, string)
	generator = func(left, right int, str string) {
		if left == 0 && right == 0 {
			result = append(result, str)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			generator(left-1, right, str+"(")
		}
		if right > 0 {
			generator(left, right-1, str+")")
		}
	}
	generator(n, n, "")
	return result
}

func five(n int) []string {
	var result []string
	var handler func(int, int, string)
	handler = func(left int, right int, s string) {
		if left == 0 && right == 0 {
			result = append(result, s)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			handler(left-1, right, s+"(")
		}
		if right > 0 {
			handler(left, right-1, s+")")
		}
	}
	handler(n, n, "")
	return result
}

func six(n int) []string {
	var result []string
	var handler func(int, int, string)
	handler = func(left int, right int, s string) {
		if left == 0 && right == 0 {
			result = append(result, s)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			handler(left-1, right, s+"(")
		}
		if right > 0 {
			handler(left, right-1, s+")")
		}
	}
	handler(n, n, "")
	return result
}

func seven(n int) []string {
	var result []string
	var handler func(int, int, string)
	handler = func(left int, right int, s string) {
		if left == 0 && right == 0 {
			result = append(result, s)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			handler(left-1, right, s+"(")
		}
		if right > 0 {
			handler(left, right-1, s+")")
		}
	}
	handler(n, n, "")
	return result
}

func eight(n int) []string {
	var result []string
	var handler func(int, int, string)
	handler = func(left int, right int, s string) {
		if left == 0 && right == 0 {
			result = append(result, s)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			handler(left-1, right, s+"(")
		}
		if right > 0 {
			handler(left, right-1, s+")")
		}
	}
	handler(n, n, "")
	return result
}

func nine(n int) []string {
	var result []string
	var handler func(int, int, string)
	handler = func(left int, right int, s string) {
		if left == 0 && right == 0 {
			result = append(result, s)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			handler(left-1, right, s+"(")
		}
		if right > 0 {
			handler(left, right-1, s+")")
		}
	}
	handler(n, n, "")
	return result
}

func ten(n int) []string {
	var result []string
	var handler func(int, int, string)
	handler = func(left int, right int, str string) {
		if left == 0 && right == 0 {
			result = append(result, str)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			handler(left-1, right, str+"(")
		}
		if right > 0 {
			handler(left, right-1, str+")")
		}
	}
	handler(n, n, "")
	return result
}

func eleven(n int) []string {
	var result []string
	var handler func(int, int, string)
	handler = func(left int, right int, str string) {
		if left == 0 && right == 0 {
			result = append(result, str)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			handler(left-1, right, str+"(")
		}
		if right > 0 {
			handler(left, right-1, str+")")
		}
	}
	handler(n, n, "")
	return result
}

func main() {
	test := 3
	fmt.Println(generateParenthesis(test))
}
