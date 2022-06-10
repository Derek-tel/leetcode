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

			if mapDic[num+1] {
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
