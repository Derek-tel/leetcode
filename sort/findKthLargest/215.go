package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {

	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)

	for i := len(nums) - 1; i > len(nums)-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}

	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, heapSize int) {
	left := i*2 + 1
	right := i*2 + 2
	largest := i
	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}
	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, heapSize)
	}

}

func test(nums []int, k int) int {
	headSize := len(nums)
	build(nums, headSize)

	for i := len(nums) - 1; i > len(nums)-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		headSize--
		helper(nums, 0, headSize)
	}
	return nums[0]
}

func build(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		helper(nums, i, heapSize)
	}
}

func helper(nums []int, i int, heapSize int) {
	left := i*2 + 1
	right := i*2 + 2
	largest := i
	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}
	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		helper(nums, largest, heapSize)
	}
}

func get(nums []int, k int) int {
	size := len(nums)
	buildHeap(nums, size)
	for i := len(nums) - 1; i > len(nums)-k; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		size--
		done(nums, 0, size)
	}
	return nums[0]
}

func buildHeap(nums []int, size int) {
	for i := size / 2; i >= 0; i-- {
		done(nums, i, size)
	}
}

func done(nums []int, i int, size int) {
	left := i*2 + 1
	right := i*2 + 2
	largest := i
	if left < size && nums[left] > nums[largest] {
		largest = left
	}
	if right < size && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		done(nums, largest, size)
	}
}

func four(nums []int, k int) int {
	var buildUp func([]int)
	var modify func([]int, int, int)
	buildUp = func(ints []int) {
		size := len(ints)
		for i := size / 2; i >= 0; i-- {
			modify(ints, i, size)
		}
	}
	modify = func(ints []int, i int, size int) {
		left := i*2 + 1
		right := i*2 + 2
		largest := i
		if left < size && ints[left] > ints[largest] {
			largest = left
		}
		if right < size && ints[right] > ints[largest] {
			largest = right
		}
		if largest != i {
			ints[i], ints[largest] = ints[largest], ints[i]
			modify(ints, largest, size)
		}
	}
	buildUp(nums)
	size := len(nums)
	for i := len(nums) - 1; i > len(nums)-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		size--
		modify(nums, 0, size)
	}
	return nums[0]
}

func five(nums []int, k int) int {
	var headModify func([]int, int, int)
	headModify = func(ints []int, i int, size int) {
		left := i*2 + 1
		right := i*2 + 2
		largest := i
		if left < size && ints[left] > ints[largest] {
			largest = left
		}
		if right < size && ints[right] > ints[largest] {
			largest = right
		}
		if i != largest {
			ints[i], ints[largest] = ints[largest], ints[i]
			headModify(ints, largest, size)
		}
	}
	var buildHead func([]int, int)
	buildHead = func(ints []int, size int) {
		for i := size / 2; i >= 0; i-- {
			headModify(ints, i, size)
		}
	}
	size := len(nums)
	buildHead(nums, size)
	for i := len(nums) - 1; i > len(nums)-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		size--
		headModify(nums, 0, size)
	}
	return nums[0]
}

func main() {
	test := []int{1, 4, 2, 7, 9, -1}
	fmt.Println(get(test, 2))
}
