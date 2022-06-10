package main

import "fmt"

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

func main() {
	test := []int{1, 1, 3, 4, 5, -1, -2, 6, 1}
	fmt.Println(get(test, 0, len(test)-1))
}
