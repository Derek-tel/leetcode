package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	return x == flag || x == flag/10
}

func one(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	fmt.Println(x, flag)
	return x == flag || x == flag/10
}

func two(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	return x == flag || x == flag/10
}

func three(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	return x == flag || x == flag/10
}

func four(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	return x == flag || x == flag/10
}

func five(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}

	return x == flag || x == flag/10
}

func six(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	return x == flag || x == flag/10
}

func seven(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	return x == flag || x == flag/10
}

func eight(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	return x == flag || x == flag/10
}

func nine(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	return x == flag || x == flag/10
}

func ten(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	return x == flag || x == flag/10
}

func main() {
	x := 10101
	fmt.Println(one(x))
}
