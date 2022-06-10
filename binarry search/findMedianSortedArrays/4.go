package main

import "fmt"

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
