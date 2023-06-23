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
