package mySqrt

func mySqrt(x int) int {
	left, right := 0, x
	resp := -1
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid <= x {
			resp = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return resp
}

func one(x int) int {
	left, right := 0, x
	resp := -1
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid <= x {
			resp = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return resp
}
