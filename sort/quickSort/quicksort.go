package quickSort

import (
	"fmt"
)

func quickSort(value []int, start int, end int) []int {
	temp := value[start]
	p := start
	i, j := start, end

	for i < j {
		for i < j && value[j] >= temp {
			j--
		}
		if j >= p {
			value[p] = value[j]
			p = j
		}
		for i < j && value[i] < temp {
			i++
		}
		if i < p {
			value[p] = value[i]
			p = i
		}
	}
	value[p] = temp
	if p-start > 1 {
		quickSort(value, 0, p-1)
	}
	if end-p > 1 {
		quickSort(value, p+1, end)
	}
	return value
}

func get(value []int, start, end int) []int {
	p := start
	temp := value[p]
	i, j := start, end

	for i <= j {
		for j >= p && value[j] >= temp {
			j--
		}
		if j >= p {
			value[j], value[p] = value[p], value[j]
			p = j
		}

		for i <= p && value[i] < temp {
			i++
		}
		if i <= p {
			value[i], value[p] = value[p], value[i]
			p = i
		}
	}
	value[p] = temp
	if p-start > 1 {
		get(value, start, p-1)
	}
	if end-p > 1 {
		get(value, p+1, end)
	}
	return value
}

func demo(nums []int, start, end int) []int {
	p := start
	temp := nums[p]
	i, j := start, end
	for i <= j {
		for j >= p && nums[j] >= temp {
			j--
		}
		if j >= p {
			nums[j], nums[p] = nums[p], nums[j]
			p = j
		}
		for i <= p && nums[i] < temp {
			i++
		}
		if i <= p {
			nums[i], nums[p] = nums[p], nums[i]
			p = i
		}
	}
	nums[p] = temp
	if p-start > 1 {
		demo(nums, start, p-1)
	}
	if end-p > 1 {
		demo(nums, p+1, end)
	}
	return nums
}

func four(nums []int) []int {
	var sort func([]int, int, int)
	sort = func(ints []int, left int, right int) {
		if left >= right {
			return
		}
		start, end := left, right
		nums[start], nums[(start+end)>>1] = nums[(start+end)>>1], nums[start]
		pivot := nums[start]
		for start < end {
			for start < end && pivot <= nums[end] {
				end--
			}
			nums[start] = nums[end]
			for start < end && pivot >= nums[start] {
				start++
			}
			nums[end] = nums[start]
		}
		nums[start] = pivot
		sort(nums, left, start-1)
		sort(nums, start+1, right)
	}

	sort(nums, 0, len(nums)-1)
	return nums
}

func five(nums []int) []int {
	var sort func([]int, int, int)
	sort = func(ints []int, left int, right int) {
		if left >= right {
			return
		}
		start, end := left, right
		ints[start], ints[(start+end)>>1] = ints[(start+end)>>1], ints[start]
		pivot := ints[start]
		for start < end {
			for start < end && pivot <= nums[end] {
				end--
			}
			ints[start] = ints[end]
			for start < end && pivot >= nums[start] {
				start++
			}
			ints[end] = ints[start]
		}
		ints[start] = pivot
		sort(ints, left, start-1)
		sort(ints, start+1, right)
	}

	sort(nums, 0, len(nums)-1)
	return nums
}

func six(nums []int) []int {
	var sort func([]int, int, int)
	sort = func(ints []int, left int, right int) {
		if left >= right {
			return
		}
		start, end := left, right
		ints[start], ints[(start+end)>>1] = ints[(start+end)>>1], ints[start]
		p := ints[start]
		for start < end {
			for start < end && p <= ints[end] {
				end--
			}
			ints[start] = ints[end]
			for start < end && p >= ints[start] {
				start++
			}
			ints[end] = ints[start]
		}
		ints[start] = p
		sort(ints, left, start-1)
		sort(ints, start+1, right)
	}
	sort(nums, 0, len(nums)-1)
	return nums
}

func seven(nums []int, start, end int) []int {
	p := start
	temp := nums[start]
	i, j := start, end
	for i <= j {
		for j >= p && nums[j] >= temp {
			j--
		}
		if j >= p {
			nums[j], nums[p] = nums[p], nums[j]
			p = j
		}
		for i <= p && nums[i] <= p {
			i++
		}
		if i <= p {
			nums[i], nums[p] = nums[p], nums[i]
			p = i
		}
	}
	nums[p] = temp
	if p-start >= 1 {
		seven(nums, start, p-1)
	}
	if end-p >= 1 {
		seven(nums, p+1, end)
	}
	return nums
}

func eight(nums []int, start, end int) []int {
	pivot := start
	temp := nums[pivot]
	i, j := start, end
	for i <= j {
		for j >= pivot && nums[j] >= temp {
			j--
		}
		if j >= pivot {
			nums[j], nums[pivot] = nums[pivot], nums[j]
			pivot = j
		}
		for i <= pivot && nums[i] < temp {
			i++
		}
		if i <= pivot {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot = i
		}
	}
	nums[pivot] = temp
	if pivot-start > 1 {
		eight(nums, start, pivot-1)
	}
	if end-pivot > 1 {
		eight(nums, pivot+1, end)
	}
	return nums
}

func nine(nums []int, start, end int) []int {
	pivot := start
	temp := nums[pivot]
	i, j := start, end
	for i <= j {
		for j >= pivot && nums[j] >= temp {
			j--
		}
		if j >= pivot {
			nums[j], nums[pivot] = nums[pivot], nums[j]
			pivot = j
		}
		for i <= pivot && nums[i] < temp {
			i++
		}
		if i <= pivot {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot = i
		}
	}
	nums[pivot] = temp
	if pivot-start > 1 {
		nine(nums, start, pivot-1)
	}
	if end-pivot > 1 {
		nine(nums, pivot+1, end)
	}
	return nums
}

func ten(nums []int, start, end int) []int {
	pivot := start
	temp := nums[pivot]
	i, j := start, end
	for i <= j {
		for j >= pivot && nums[j] > temp {
			j--
		}
		if j >= pivot {
			nums[j], nums[pivot] = nums[pivot], nums[j]
			pivot = j
		}
		for i <= pivot && nums[i] < temp {
			i++
		}
		if i <= pivot {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot = i
		}
	}
	nums[pivot] = temp
	if pivot-start > 1 {
		ten(nums, start, pivot-1)
	}
	if end-pivot > 1 {
		ten(nums, pivot+1, end)
	}
	return nums
}

func eleven(nums []int, start, end int) []int {
	pivot := start
	temp := nums[pivot]
	i, j := start, end
	for i <= j {
		for j >= pivot && nums[j] > temp {
			j--
		}
		if j >= pivot {
			nums[j], nums[pivot] = nums[pivot], nums[j]
			pivot = j
		}
		for i <= pivot && nums[i] < temp {
			i++
		}
		if i <= pivot {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot = i
		}
	}
	nums[pivot] = temp
	if pivot-start > 1 {
		eleven(nums, start, pivot-1)
	}
	if end-pivot > 1 {
		eleven(nums, pivot+1, end)
	}
	return nums
}

func twelve(nums []int, start, end int) []int {
	pivot := start
	temp := nums[pivot]
	i, j := start, end
	for i <= j {
		for j >= pivot && nums[j] > temp {
			j--
		}
		if j >= pivot {
			nums[j], nums[pivot] = nums[pivot], nums[j]
			pivot = j
		}
		for i <= pivot && nums[i] < temp {
			i++
		}
		if i <= pivot {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot = i
		}
	}
	nums[pivot] = temp
	if pivot-start > 1 {
		twelve(nums, start, pivot-1)
	}
	if end-pivot > 1 {
		twelve(nums, pivot+1, end)
	}
	return nums
}

func thirteen(nums []int, start, end int) []int {
	pivot := start
	temp := nums[pivot]
	i, j := start, end
	for i <= j {
		for j >= pivot && nums[j] > temp {
			j--
		}
		if j >= pivot {
			nums[j], nums[pivot] = nums[pivot], nums[j]
			pivot = j
		}
		for i <= pivot && nums[i] < temp {
			i++
		}
		if i <= pivot {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot = i
		}
	}
	nums[pivot] = temp

	if pivot-start > 1 {
		thirteen(nums, start, pivot-1)
	}
	if end-pivot > 1 {
		thirteen(nums, pivot+1, end)
	}
	return nums
}

func fourteen(nums []int, start, end int) []int {
	pivot := start
	temp := nums[pivot]
	i, j := start, end
	for i <= j {
		for j >= pivot && nums[j] >= temp {
			j--
		}
		if j >= pivot {
			nums[j], nums[pivot] = nums[pivot], nums[j]
			pivot = j
		}
		for i <= pivot && nums[i] < temp {
			i++
		}
		if i <= pivot {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot = i
		}
	}
	nums[pivot] = temp
	if pivot-start > 1 {
		fourteen(nums, start, pivot-1)
	}

	if end-pivot > 1 {
		fourteen(nums, pivot+1, end)
	}
	return nums
}

func main() {
	test := []int{1, 1, 3, 4, 5, -1, -2, 6, 1}
	fmt.Println(four(test))
}
