package main

import "fmt"

func one(n int) []int {
	ans := make([]int, n)
	num := 1
	for i := range ans {
		ans[i] = num
		if num*10 <= n {
			num = num * 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return ans
}

func two(n int) []int {
	var resp []int
	var handler func(int)
	handler = func(index int) {
		if index > n {
			return
		}
		resp = append(resp, index)
		for i := 0; i <= 9; i++ {
			handler(index*10 + i)
		}
	}
	for i := 1; i <= 9; i++ {
		handler(i)
	}
	return resp
}

func three(n int) []int {
	var resp []int
	var handler func(int)
	handler = func(index int) {
		if index > n {
			return
		}
		resp = append(resp, index)
		for i := 0; i <= 9; i++ {
			handler(index*10 + i)
		}
	}
	for i := 1; i <= 9; i++ {
		handler(i)
	}
	return resp
}

func four(n int) []int {
	var resp []int
	var handler func(int)
	handler = func(index int) {
		if index > n {
			return
		}
		resp = append(resp, index)
		for i := 0; i <= 9; i++ {
			handler(index*10 + i)
		}
	}
	for i := 1; i <= 9; i++ {
		handler(i)
	}
	return resp
}

func five(n int) []int {
	var resp []int
	var handler func(int)
	handler = func(index int) {
		if index > n {
			return
		}
		resp = append(resp, index)
		for i := 0; i <= 9; i++ {
			handler(index*10 + i)
		}
	}
	for i := 1; i <= 9; i++ {
		handler(i)
	}
	return resp
}

func six(n int) []int {
	var resp []int
	var handler func(int)
	handler = func(index int) {
		if index > n {
			return
		}
		resp = append(resp, index)
		for i := 0; i <= 9; i++ {
			handler(index*10 + i)
		}
	}

	for i := 1; i <= 9; i++ {
		handler(i)
	}
	return resp
}

func main() {
	fmt.Println(two(13))
}
