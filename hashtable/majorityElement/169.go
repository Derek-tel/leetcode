package majorityElement

func majorityElement(nums []int) int {
	count := 0
	var p int

	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if num != p {
				count--
			} else {
				count++
			}
		}
	}
	return p
}

func test(nums []int) int {
	count := 0

	var p int

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			p = nums[i]
			count++
		} else {
			if nums[i] == p {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func get(nums []int) int {
	count := 0
	var p int

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			p = nums[i]
			count++
		} else {
			if nums[i] == p {
				count++
			} else {
				count--
			}
		}
	}
	return p
}
