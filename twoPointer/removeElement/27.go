package removeElement

func removeElement(nums []int, val int) int {
	flag := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[flag], nums[i] = nums[i], nums[flag]
			flag++
		}
	}
	return flag
}

func first(nums []int, val int) int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[flag], nums[i] = nums[i], nums[flag]
			flag++
		}
	}

	return flag
}

func two(nums []int, val int) int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[flag], nums[i] = nums[i], nums[flag]
			flag++
		}
	}
	return flag
}

func three(nums []int, val int) int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[flag], nums[i] = nums[i], nums[flag]
			flag++
		}
	}
	return flag
}

func four(nums []int, val int) int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[flag], nums[i] = nums[i], nums[flag]
			flag++
		}
	}
	return flag
}

func five(nums []int, val int) int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return flag
}

func six(nums []int, val int) int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return flag
}

func seven(nums []int, val int) int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return flag
}

func eight(nums []int, val int) int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	return flag
}
