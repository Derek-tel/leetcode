package searchInsert

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return len(nums)
}

func one(num []int, target int) int {
	left, right := 0, len(num)-1
	for left <= right {
		mid := left + (right-left)>>1
		if num[mid] >= target {
			if mid == 0 || num[mid-1] < target {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return len(num)
}
