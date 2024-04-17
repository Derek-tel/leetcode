package sortArray

import (
	"math/rand"
)

func one(nums []int) []int {
	var handler func([]int, int, int) []int
	handler = func(ints []int, start, end int) []int {
		count := end - start
		pivot := start
		if count > 0 {
			pivot = rand.Intn(count) + start
		}
		temp := ints[pivot]
		i, j := start, end
		for i <= j {
			for j >= pivot && ints[j] >= temp {
				j--
			}
			if j >= pivot {
				ints[j], ints[pivot] = ints[pivot], ints[j]
				pivot = j
			}
			for i <= pivot && ints[i] <= temp {
				i++
			}
			if i <= pivot {
				ints[i], ints[pivot] = ints[pivot], ints[i]
				pivot = i
			}
		}
		ints[pivot] = temp
		if pivot-1 > start {
			handler(ints, start, pivot-1)
		}
		if pivot+1 < end {
			handler(ints, pivot+1, end)
		}
		return ints
	}
	handler(nums, 0, len(nums)-1)
	return nums
}

func two(nums []int) []int {
	var handler func([]int, int, int) []int
	handler = func(ints []int, start int, end int) []int {
		count := end - start
		pivot := start
		if count > 0 {
			pivot = start + rand.Intn(count)
		}
		temp := ints[pivot]
		i, j := start, end
		for i <= j {
			for pivot <= j && ints[j] >= temp {
				j--
			}
			if pivot <= j {
				ints[j], ints[pivot] = ints[pivot], ints[j]
				pivot = j
			}
			for i <= pivot && ints[i] < temp {
				i++
			}
			if i <= pivot {
				ints[i], ints[pivot] = ints[pivot], ints[i]
				pivot = i
			}
		}
		ints[pivot] = temp
		if start < pivot-1 {
			handler(ints, start, pivot-1)
		}
		if pivot+1 < end {
			handler(ints, pivot+1, end)
		}
		return ints
	}
	handler(nums, 0, len(nums)-1)
	return nums
}

func three(nums []int) []int {
	var handler func([]int, int, int)
	handler = func(ints []int, start int, end int) {
		count := end - start
		pivot := start
		if count > 0 {
			pivot = start + rand.Intn(count)
		}
		temp := ints[pivot]
		i, j := start, end
		for i <= j {
			for pivot <= j && ints[j] >= temp {
				j--
			}
			if pivot <= j {
				ints[j], ints[pivot] = ints[pivot], ints[j]
				pivot = j
			}
			for i <= pivot && ints[i] < temp {
				i++
			}
			if i <= pivot {
				ints[i], ints[pivot] = ints[pivot], ints[i]
				pivot = i
			}
		}
		ints[pivot] = temp
		if pivot-1 > start {
			handler(ints, start, pivot-1)
		}
		if pivot+1 < end {
			handler(ints, pivot+1, end)
		}
	}

	handler(nums, 0, len(nums)-1)
	return nums
}
