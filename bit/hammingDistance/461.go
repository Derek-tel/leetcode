package hammingDistance

func hammingDistance(x int, y int) int {
	ans := 0

	for i := x ^ y; i > 0; i = i >> 1 {
		ans = ans + i&1
	}
	return ans
}

func test(x int, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans = ans + i&1
	}
	return ans
}

func get(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans += i & 1
	}
	return ans
}

func four(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans += i & 1
	}
	return ans
}

func five(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans += i & 1
	}
	return ans
}

func six(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; {
		i = i & (i - 1)
		ans++
	}
	return ans
}

func seven(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans = ans + i&1
	}
	return ans
}

func eight(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans = ans + i&1
	}
	return ans
}

func nine(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans = ans + i&1
	}
	return ans
}

func ten(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans = ans + i&1
	}
	return ans
}

func eleven(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans += i & 1
	}
	return ans
}

func twelve(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans += i & 1
	}
	return ans
}

func thirteen(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans += i & 1
	}
	return ans
}

func fourteen(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans += 1 & i
	}
	return ans
}

func fifteen(x, y int) int {
	ans := 0
	for i := x ^ y; i > 0; i = i >> 1 {
		ans += i & 1
	}
	return ans
}
