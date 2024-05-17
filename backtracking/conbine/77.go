package conbine

func one(n int, k int) [][]int {
	var temp []int
	var result [][]int
	var handler func(int)
	handler = func(index int) {
		if len(temp) == k {
			result = append(result, append([]int(nil), temp...))
			return
		}
		if index > n {
			return
		}

		//no
		handler(index + 1)

		//yes
		temp = append(temp, index)
		handler(index + 1)
		temp = temp[:len(temp)-1]
	}
	handler(1)
	return result
}
