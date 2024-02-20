package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left, right := 0, length-1

	for left <= right {
		half := (left + right) / 2
		if nums[half] == target {
			return half
		}
		if nums[0] <= nums[half] {
			if nums[0] <= target && target < nums[half] {
				right = half - 1
			} else {
				left = half + 1
			}
		} else {
			if nums[half] < target && nums[length-1] >= target {
				left = half + 1
			} else {
				right = half - 1
			}
		}
	}
	return -1
}

func test(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		half := (right + left) / 2
		if nums[half] == target {
			return half
		}
		if nums[0] <= nums[half] {
			if target >= nums[0] && nums[half] > target {
				right = half - 1
			} else {
				left = half + 1
			}
		} else {
			if nums[half] < target && target <= nums[right] {
				left = half + 1
			} else {
				right = half - 1
			}
		}
	}
	return -1
}

func get(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		half := (left + right) / 2
		if nums[half] == target {
			return half
		}
		if nums[0] <= nums[half] {
			if target >= nums[0] && nums[half] > target {
				right = half - 1
			} else {
				left = half + 1
			}
		} else {
			if nums[half] < target && target <= nums[length-1] {
				left = half + 1
			} else {
				right = half - 1
			}
		}
	}
	return -1
}

func four(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		half := left + (right-left)>>1
		if nums[half] == target {
			return half
		}
		if nums[0] <= nums[half] {
			if target >= nums[0] && nums[half] > target {
				right = half - 1
			} else {
				left = half + 1
			}
		} else {
			if nums[half] < target && target <= nums[length-1] {
				left = half + 1
			} else {
				right = half - 1
			}
		}
	}
	return -1
}

func five(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		half := left + (right-left)>>1
		if nums[half] == target {
			return half
		}
		if nums[0] <= nums[half] {
			if target >= nums[0] && nums[half] > target {
				right = half - 1
			} else {
				left = half + 1
			}
		} else {
			if nums[half] < target && target <= nums[length-1] {
				left = half + 1
			} else {
				right = half - 1
			}
		}
	}
	return -1
}

func six(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		half := left + (right-left)>>1
		if nums[half] == target {
			return half
		}
		if nums[0] <= nums[half] {
			if target >= nums[0] && nums[half] > target {
				right = half - 1
			} else {
				left = half + 1
			}
		} else {
			if nums[half] < target && target <= nums[length-1] {
				left = half + 1
			} else {
				right = half - 1
			}
		}
	}
	return -1
}

func seven(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[length-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func eight(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && nums[length-1] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func nine(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[length-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func ten(nums []int, target int) int {
	length := len(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[length-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func eleven(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[length-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func twelve(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[length-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func thirteen(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] == nums[left] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func fourteen(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] == nums[left] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	test := []int{3, 1}
	fmt.Println(search(test, 1))
}
