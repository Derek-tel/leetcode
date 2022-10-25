package reverse

import "math"

func reverse(x int) int {
	var result int
	for x != 0 {
		if result < math.MinInt32/10 || result > math.MaxInt32/10 {
			return 0
		}
		digit := x & 10
		x = x / 10
		result = result*10 + digit
	}
	return result
}
