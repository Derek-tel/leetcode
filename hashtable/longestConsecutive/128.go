package main

func longestConsecutive(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	mapDict := make(map[int]bool)
	for _, num := range nums {
		mapDict[num] = true
	}
	longest := 0
	for _, num := range nums {
		if !mapDict[num-1] {
			current := 1

			for mapDict[num+1] {
				current++
				num++
			}

			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func get(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	mapDict := make(map[int]bool)
	for _, num := range nums {
		mapDict[num] = true
	}
	longest := 0
	for _, num := range nums {
		if !mapDict[num-1] {
			currentLong := 1
			for mapDict[num+1] {
				currentLong++
				num++
			}
			if currentLong > longest {
				longest = currentLong
			}
		}
	}
	return longest
}

func test(nums []int) int {
	mapDic := make(map[int]bool)

	for _, num := range nums {
		mapDic[num] = true
	}

	longest := 0

	for _, num := range nums {
		if !mapDic[num-1] {
			current := 1

			for mapDic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func four(nums []int) int {
	dic := make(map[int]bool)

	for _, n := range nums {
		dic[n] = true
	}
	longest := 0
	for _, num := range nums {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func five(nums []int) int {
	dic := make(map[int]bool)
	for _, num := range nums {
		dic[num] = true
	}
	longest := 0
	for _, num := range nums {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func six(nums []int) int {
	dic := make(map[int]bool)
	for _, v := range nums {
		dic[v] = true
	}
	longest := 0
	for _, num := range nums {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}
func seven(nums []int) int {
	dic := make(map[int]bool)
	for _, n := range nums {
		dic[n] = true
	}
	longest := 0
	for _, num := range nums {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func eight(nums []int) int {
	dic := make(map[int]bool)
	for _, n := range nums {
		dic[n] = true
	}
	longest := 0
	for _, num := range nums {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func nine(nums []int) int {
	dic := make(map[int]bool)
	for _, num := range nums {
		dic[num] = true
	}
	longest := 0
	for _, num := range nums {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func ten(nums []int) int {
	dic := make(map[int]bool)
	for _, num := range nums {
		dic[num] = true
	}
	longest := 0
	for num := range dic {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func eleven(nums []int) int {
	dic := make(map[int]bool)
	for _, num := range nums {
		dic[num] = true
	}
	longest := 0
	for num := range dic {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func twelve(nums []int) int {
	dic := make(map[int]bool)
	for _, num := range nums {
		dic[num] = true
	}
	longest := 0
	for num := range dic {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func thirteen(nums []int) int {
	dic := make(map[int]bool)
	for _, num := range nums {
		dic[num] = true
	}
	longest := 0
	for num := range dic {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func fourteen(nums []int) int {
	dic := make(map[int]bool)
	for _, num := range nums {
		dic[num] = true
	}
	longest := 0
	for num := range dic {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func fifteen(nums []int) int {
	dic := make(map[int]bool)
	for _, num := range nums {
		dic[num] = true
	}
	longest := 0
	for num := range dic {
		if !dic[num-1] {
			current := 1
			for dic[num+1] {
				current++
				num++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}
