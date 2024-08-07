package main

import (
	"fmt"
)

func countBits(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func countBit(n int) []int {
	bit := make([]int, n+1)
	flag := 0
	bit[0] = 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			flag = i
		}
		bit[i] = 1 + bit[i-flag]
	}
	return bit
}

func get(n int) []int {
	bit := make([]int, n+1)
	high := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			high = i
		}
		bit[i] = bit[i-high] + 1
	}
	return bit
}

func four(n int) []int {
	bit := make([]int, n+1)
	high := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			high = i
		}
		bit[i] = bit[i-high] + 1
	}
	return bit
}

func five(n int) []int {
	bit := make([]int, n+1)
	high := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			high = i
		}
		bit[i] = bit[i-high] + 1
	}
	return bit
}

func six(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func seven(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func eight(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func nine(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func ten(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func eleven(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func twelve(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func thirteen(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func fourteen(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func fifteen(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func sixteen(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func seventeen(n int) []int {
	bit := make([]int, n+1)
	highBit := 0
	bit[0] = 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bit[i] = bit[i-highBit] + 1
	}
	return bit
}

func main() {
	fmt.Println(countBits(5))
}
