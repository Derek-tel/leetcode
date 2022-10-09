package main

import (
	"fmt"
)

func HeapSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}
	//将待排序的数组划分为左右2部分，递归的进行
	mid := length / 2
	left := data[:mid]
	right := data[mid:]
	left = HeapSort(left)
	right = HeapSort(right)
	//对划分后的两部分进行排序
	return Sort(left, right)
}

func Sort(left []int, right []int) []int {
	llen := len(left)
	rlen := len(right)
	var tmp []int

	i, j := 0, 0

	for {
		//如果left先遍历结束，则将right剩余部分追加到tmp中
		if i >= llen {
			tmp = append(tmp, right[j:]...)
			break
		}
		//如果right部分先遍历结束，则将left剩余部分追加到tmp中
		if j >= rlen {
			tmp = append(tmp, left[i:]...)
			break
		}
		//比较left和right将较小的元素存入tmp
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	return tmp
}

func get(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	mid := length / 2
	left := nums[:mid]
	right := nums[mid:]
	return merge(get(left), get(right))
}

func merge(left, right []int) []int {
	llen := len(left)
	rlen := len(right)
	var temp []int

	i, j := 0, 0

	for {
		if i >= llen {
			temp = append(temp, left[i:]...)
			break
		}
		if j >= rlen {
			temp = append(temp, right[j:]...)
			break
		}
		if left[i] < right[j] {
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			j++
		}
	}
	return temp
}

func test(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	mid := length / 2
	left := test(nums[:mid])
	right := test(nums[mid:])
	return helper(left, right)
}

func helper(left, right []int) []int {
	lengthA := len(left)
	lengthB := len(right)
	i, j := 0, 0
	var res []int
	for {
		if i >= lengthA {
			res = append(res, right[j:]...)
			break
		}
		if j >= lengthB {
			res = append(res, left[i:]...)
			break
		}
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	return res
}

func four(nums []int) []int {
	var handler func([]int, []int) []int
	handler = func(left []int, right []int) []int {
		var temp []int
		i, j := 0, 0
		for {
			if i >= len(left) {
				temp = append(temp, right[j:]...)
				break
			}
			if j >= len(right) {
				temp = append(temp, left[i:]...)
				break
			}
			if left[i] < right[j] {
				temp = append(temp, left[i])
				i++
			} else {
				temp = append(temp, right[j])
				j++
			}
		}
		return temp
	}
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]
	return handler(four(left), four(right))
}

// 测试：
func main() {
	arr := []int{0, 1, 5, 2, 3, 8, 0, 4, 9, 2}
	r := four(arr)
	fmt.Println(r)
}
