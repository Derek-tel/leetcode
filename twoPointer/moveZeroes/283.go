package moveZeroes

func moveZeroes(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	flag := 0

	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return nums
}

func test(nums []int) []int {
	length := len(nums)
	flag := 0

	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return nums
}

func get(nums []int) []int {
	length := len(nums)
	flag := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return nums
}

func four(nums []int) []int {
	length := len(nums)
	flag := 0

	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return nums
}

func five(nums []int) []int {
	length := len(nums)
	flag := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return nums
}

func six(nums []int) []int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return nums
}

func seven(nums []int) []int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return nums
}

func eight(nums []int) []int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return nums
}
