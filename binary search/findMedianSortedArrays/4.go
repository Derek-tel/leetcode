package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lengthA := len(nums1)
	lengthB := len(nums2)
	count := lengthA + lengthB
	if count%2 == 1 {
		K := count/2 + 1
		return float64(getK(nums1, nums2, K))
	} else {
		K := count / 2
		return float64(getK(nums1, nums2, K)+getK(nums1, nums2, K+1)) / 2
	}

}

func get(nums1, nums2 []int, k int) int {
	indexA := 0
	indexB := 0

	for {
		if indexA == len(nums1) {
			return nums2[indexB+k-1]
		}
		if indexB == len(nums2) {
			return nums1[indexA+k-1]
		}
		if k == 1 {
			return min(nums1[indexA], nums2[indexB])
		}
		half := k / 2
		newIndexA := min(len(nums1), indexA+half) - 1
		newIndexB := min(len(nums2), indexB+half) - 1
		if nums1[newIndexA] < nums2[newIndexB] {
			k = k - (newIndexA - indexA + 1)
			indexA = newIndexA + 1
		} else {
			k = k - (newIndexB - indexB + 1)
			indexB = newIndexB + 1
		}
	}
}

func getK(nums1 []int, nums2 []int, k int) int {
	indexA := 0
	indexB := 0

	for {
		if indexA == len(nums1) { //length =2
			return nums2[indexB+k-1]
		}
		if indexB == len(nums2) {
			return nums1[indexA+k-1]
		}

		if k == 1 {
			return min(nums1[indexA], nums2[indexB])
		}

		half := k / 2 //==3
		newIndexA := min(len(nums1), indexA+half) - 1
		newIndexB := min(len(nums2), indexB+half) - 1
		if nums1[newIndexA] < nums2[newIndexB] {
			k = k - (newIndexA - indexA + 1)
			indexA = newIndexA + 1
		} else {
			k = k - (newIndexB - indexB + 1)
			indexB = newIndexB + 1
		}

	}

}

func test(nums1 []int, nums2 []int, k int) int {
	indexA := 0
	indexB := 0
	for {
		if indexA == len(nums1) {
			return nums2[indexB+k-1]
		}
		if indexB == len(nums2) {
			return nums1[indexA+k-1]
		}
		if k == 1 {
			return min(nums1[indexA], nums2[indexB])
		}
		half := k / 2
		newIndexA := min(len(nums1), indexA+half) - 1
		newIndexB := min(len(nums2), indexB+half) - 1
		if nums1[newIndexA] < nums2[newIndexB] {
			k = k - (newIndexA - indexA + 1)
			indexA = newIndexA + 1
		} else {
			k = k - (newIndexB - indexB + 1)
			indexB = newIndexB + 1
		}

	}
}

func four(nums1 []int, nums2 []int) float64 {
	lengthA := len(nums1)
	lengthB := len(nums2)
	var helper func([]int, []int, int) int
	helper = func(ints1 []int, ints2 []int, i int) int {
		indexA := 0
		indexB := 0
		for {
			if indexA == len(ints1) {
				return ints2[indexB+i-1]
			}
			if indexB == len(ints2) {
				return ints1[indexA+i-1]
			}
			if i == 1 {
				return min(ints1[indexA], ints2[indexB])
			}
			half := i / 2
			newIndexA := min(len(ints1), indexA+half) - 1
			newIndexB := min(len(ints2), indexB+half) - 1
			if ints1[newIndexA] < ints2[newIndexB] {
				i = i - (newIndexA - indexA + 1)
				indexA = newIndexA + 1
			} else {
				i = i - (newIndexB - indexB + 1)
				indexB = newIndexB + 1
			}
		}
	}
	count := lengthA + lengthB
	if count%2 == 1 {
		k := count/2 + 1
		return float64(helper(nums1, nums2, k))
	} else {
		k := count / 2
		return float64(helper(nums1, nums2, k)+helper(nums1, nums2, k+1)) / 2
	}
}

func five(nums1 []int, nums2 []int) float64 {
	var getMin func(int, int) int
	getMin = func(i int, j int) int {
		if i < j {
			return i
		}
		return j
	}
	var helper func(int) int
	helper = func(i int) int {
		indexA, indexB := 0, 0
		for {
			if indexA == len(nums1) {
				return nums2[indexB+i-1]
			}
			if indexB == len(nums2) {
				return nums1[indexA+i-1]
			}
			if i == 1 {
				return getMin(nums1[indexA], nums2[indexB])
			}
			half := i / 2
			newIndexA := getMin(len(nums1), indexA+half) - 1
			newIndexB := getMin(len(nums2), indexB+half) - 1
			if nums1[newIndexA] < nums2[newIndexB] {
				i = i - (newIndexA - indexA + 1)
				indexA = newIndexA + 1
			} else {
				i = i - (newIndexB - indexB + 1)
				indexB = newIndexB + 1
			}
		}
	}
	lengthA, lengthB := len(nums1), len(nums2)
	total := lengthA + lengthB
	if total%2 == 1 {
		k := total/2 + 1
		return float64(helper(k))
	} else {
		k := total / 2
		return float64(helper(k)+helper(k+1)) / 2
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	testA := []int{1, 3}
	testB := []int{2}
	fmt.Println(findMedianSortedArrays(testA, testB))
}
