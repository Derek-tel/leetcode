package main

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	//length := len(nums)
	for i, v := range nums {
		if i <= farthest {
			farthest = max(farthest, i+v)
		} else {
			return false
		}
	}
	return true
}

func test(nuns []int) bool {
	length := len(nuns)
	if length == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < length; i++ {
		if i <= farthest {
			farthest = max(farthest, i+nuns[i])
		} else {
			return false
		}
	}
	return true
}

func demo(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < length; i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func four(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < length; i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func five(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < length; i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func six(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < length; i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func seven(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func eight(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func nine(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func ten(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func eleven(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func twelve(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func thirteen(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func fourteen(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func fifteen(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func sixteen(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
