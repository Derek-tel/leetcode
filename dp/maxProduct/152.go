package maxProduct

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	minNum, maxNum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mm := maxNum, minNum
		maxNum = max(mx*nums[i], max(mm*nums[i], nums[i]))
		minNum = min(mx*nums[i], min(mm*nums[i], nums[i]))
		res = max(res, maxNum)
	}
	return res
}

func test(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxN, minN, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(mm*nums[i], max(mn*nums[i], nums[i]))
		minN = min(mm*nums[i], min(mn*nums[i], nums[i]))
		res = max(res, maxN)
	}
	return res
}

func demo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxN, minN, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(max(mm*nums[i], mn*nums[i]), nums[i])
		minN = min(min(mm*nums[i], mn*nums[i]), nums[i])
		res = max(res, maxN)
	}
	return res
}

func get(nums []int) int {
	maxN, minN, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(max(mm*nums[i], mn*nums[i]), nums[i])
		minN = min(min(mm*nums[i], mn*nums[i]), nums[i])
		res = max(maxN, res)
	}
	return res
}

func five(nums []int) int {
	maxN, minN, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(max(mm*nums[i], mn*nums[i]), nums[i])
		minN = min(min(mm*nums[i], mn*nums[i]), nums[i])
		result = max(maxN, result)
	}
	return result
}

func six(nums []int) int {
	maxN, minN, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(max(nums[i]*mm, nums[i]*mn), nums[i])
		minN = min(min(nums[i]*mm, nums[i]*mn), nums[i])
		result = max(result, maxN)
	}
	return result
}

func seven(nums []int) int {
	maxN, minN, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(max(nums[i]*mm, nums[i]*mn), nums[i])
		minN = min(min(nums[i]*mm, nums[i]*mn), nums[i])
		result = max(result, maxN)
	}
	return result
}

func eight(nums []int) int {
	maxN, minN, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(max(nums[i]*mn, nums[i]*mm), nums[i])
		minN = min(min(nums[i]*mn, nums[i]*mm), nums[i])
		result = max(result, maxN)
	}
	return result
}

func nine(nums []int) int {
	maxN, minN, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(max(nums[i]*mm, nums[i]*mn), nums[i])
		minN = min(min(nums[i]*mm, nums[i]*mn), nums[i])
		result = max(result, maxN)
	}
	return result
}

func ten(nums []int) int {
	maxN, minN, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(max(nums[i]*mm, nums[i]*mn), nums[i])
		minN = min(min(nums[i]*mm, nums[i]*mn), nums[i])
		result = max(result, maxN)
	}
	return result
}

func eleven(nums []int) int {
	maxN, minN, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(max(nums[i]*mm, nums[i]*mn), nums[i])
		minN = min(min(nums[i]*mm, nums[i]*mn), nums[i])
		result = max(result, maxN)
	}
	return result
}

func twelve(nums []int) int {
	maxN, minN, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mm := maxN
		mn := minN
		maxN = max(max(nums[i]*mm, nums[i]*mn), nums[i])
		minN = min(min(nums[i]*mm, nums[i]*mn), nums[i])
		result = max(result, maxN)
	}
	return result
}

func thirteen(nums []int) int {
	maxN, minN, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mm, mn := maxN, minN
		maxN = max(max(nums[i]*mm, nums[i]*mn), nums[i])
		minN = min(min(nums[i]*mm, nums[i]*mn), nums[i])
		result = max(result, maxN)
	}
	return result
}

func fourteen(nums []int) int {
	maxN, minN, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mm, mn := maxN, minN
		maxN = max(max(nums[i]*mm, nums[i]*mn), nums[i])
		minN = min(min(nums[i]*mm, nums[i]*mn), nums[i])
		result = max(result, maxN)
	}
	return result
}

func fifteen(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	maxN, minN, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mm, mn := maxN, minN
		maxN = max(max(nums[i]*mm, nums[i]*mn), nums[i])
		minN = min(min(nums[i]*mm, nums[i]*mn), nums[i])
		result = max(maxN, minN)
	}
	return result
}

func sixteen(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	maxN, minN, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mm, mn := maxN, minN
		maxN = max(max(nums[i]*mm, nums[i]*mn), nums[i])
		minN = min(min(nums[i]*mm, nums[i]*mn), nums[i])
		result = max(result, maxN)
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
