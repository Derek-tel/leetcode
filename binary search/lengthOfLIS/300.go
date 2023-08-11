package lengthOfLIS

func lengthOfLIS(nums []int) int {
	var temp []int
	for _, num := range nums {
		index := searchFirstGreaterElement(temp, num)
		if index == -1 {
			temp = append(temp, num)
		} else {
			temp[index] = num
		}
	}
	return len(temp)
}

// 二分查找第一个大于等于 target 的元素，时间复杂度 O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			if (mid == 0) || (nums[mid-1] < target) { // 找到第一个大于等于 target 的元素
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func two(nums []int) int {
	var searchFirst func([]int, int) int
	searchFirst = func(ints []int, i int) int {
		low, high := 0, len(ints)-1
		for low <= high {
			mid := low + (high-low)>>1
			if ints[mid] >= i {
				if mid == 0 || ints[mid-1] < i {
					return mid
				}
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return -1
	}

	var temp []int

	for i := 0; i < len(nums); i++ {
		index := searchFirst(temp, nums[i])
		if index == -1 {
			temp = append(temp, nums[i])
		} else {
			temp[index] = nums[i]
		}
	}
	return len(temp)
}

func three(nums []int) int {
	var searchFirst func([]int, int) int
	searchFirst = func(ints []int, i int) int {
		low, high := 0, len(ints)-1
		for low <= high {
			mid := low + (high-low)>>1
			if ints[mid] >= i {
				if mid == 0 || ints[mid-1] < i {
					return mid
				}
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return -1
	}
	var temp []int
	for _, num := range nums {
		index := searchFirst(temp, num)
		if index == -1 {
			temp = append(temp, num)
		} else {
			temp[index] = num
		}
	}
	return len(temp)
}

func four(nums []int) int {
	var searchFirst func([]int, int) int
	searchFirst = func(ints []int, i int) int {
		low, high := 0, len(ints)-1
		for low <= high {
			mid := low + (high-low)>>1
			if ints[mid] >= i {
				if mid == 0 || ints[mid-1] < i {
					return mid
				}
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return -1
	}
	var temp []int
	for _, num := range nums {
		index := searchFirst(temp, num)
		if index == -1 {
			temp = append(temp, num)
		} else {
			temp[index] = num
		}
	}
	return len(temp)
}

func five(nums []int) int {
	var searchFirst func([]int, int) int
	searchFirst = func(ints []int, i int) int {
		low, high := 0, len(ints)-1
		for low <= high {
			mid := low + (high-low)>>1
			if ints[mid] >= i {
				if mid == 0 || nums[mid-1] < i {
					return mid
				}
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return -1
	}
	var temp []int
	for _, num := range nums {
		index := searchFirst(temp, num)
		if index == -1 {
			temp = append(temp, num)
		} else {
			temp[index] = num
		}
	}
	return len(temp)
}

func six(nums []int) int {
	var searchFirst func([]int, int) int
	searchFirst = func(ints []int, i int) int {
		low, high := 0, len(ints)-1
		for low <= high {
			mid := low + (high-low)>>1
			if ints[mid] >= i {
				if mid == 0 || ints[mid-1] < i {
					return mid
				}
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return -1
	}
	var temp []int
	for _, num := range nums {
		index := searchFirst(temp, num)
		if index == -1 {
			temp = append(temp, num)
		} else {
			temp[index] = num
		}
	}
	return len(temp)
}
