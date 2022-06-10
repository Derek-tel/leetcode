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
