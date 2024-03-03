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

func five(nums []int) []int {
	var handler func([]int, []int) []int
	handler = func(sliceA []int, sliceB []int) []int {
		var temp []int
		i, j := 0, 0
		for {
			if i >= len(sliceA) {
				temp = append(temp, sliceB[j:]...)
				break
			}
			if j >= len(sliceB) {
				temp = append(temp, sliceA[i:]...)
			}
			if sliceA[i] < sliceB[j] {
				temp = append(temp, sliceA[i])
				i++
			} else {
				temp = append(temp, sliceB[j])
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
	return handler(five(left), five(right))
}

func six(nums []int) []int {
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
	left := nums[0:mid]
	right := nums[mid:]
	return handler(six(left), six(right))
}

func seven(nums []int) []int {
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
			if left[i] < right[i] {
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
	left := nums[0:mid]
	right := nums[mid:]
	return handler(seven(left), seven(right))
}

func eight(nums []int) []int {
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
	} else {
		mid := len(nums) / 2
		left := eight(nums[0:mid])
		right := eight(nums[mid:])
		return handler(left, right)
	}
}

func nine(nums []int) []int {
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
				temp = append(temp, right[i])
				j++
			}
		}
		return temp
	}

	if len(nums) <= 1 {
		return nums
	} else {
		mid := len(nums) / 2
		return handler(nine(nums[:mid]), nine(nums[mid:]))
	}
}

func ten(nums []int) []int {
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
	} else {
		mid := len(nums) / 2
		return handler(ten(nums[:mid]), ten(nums[mid:]))
	}
}

func eleven(nums []int) []int {
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
			if left[i] < right[i] {
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
	} else {
		mid := len(nums) / 2
		return handler(eleven(nums[:mid]), eleven(nums[mid:]))
	}
}

func twelve(nums []int) []int {
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
	} else {
		mid := len(nums) / 2
		return handler(twelve(nums[:mid]), twelve(nums[mid:]))
	}
}

func thirteen(nums []int) []int {
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
				temp = append(temp, right[i])
				j++
			}
		}
		return temp
	}
	if len(nums) <= 1 {
		return nums
	} else {
		mid := len(nums) / 2
		return handler(thirteen(nums[:mid]), thirteen(nums[mid:]))
	}
}

func fourteen(nums []int) []int {
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
	} else {
		mid := len(nums) / 2
		return handler(fourteen(nums[:mid]), fourteen(nums[mid:]))
	}
}

// 测试：
func main() {
	arr := []int{0, 1, 5, 2, 3, 8, 0, 4, 9, 2}
	r := ten(arr)
	fmt.Println(r)
}
