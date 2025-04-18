package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		currentA := nums1[pm]
		currentB := nums2[pn]
		if currentA >= currentB {
			nums1[p] = currentA
			pm--
		} else {
			nums1[p] = currentB
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
}

func test(nums1 []int, nums2 []int, m int, n int) []int {
	p := m + n - 1
	pm := m - 1
	pn := n - 1

	for ; pm >= 0 && pn >= 0; p-- {
		currentA := nums1[pm]
		currentB := nums2[pn]
		if currentA >= currentB {
			nums1[p] = currentA
			pm--
		} else {
			nums1[p] = currentB
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
	return nums1
}

func get(nums1, nums2 []int, m, n int) []int {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		currentA := nums1[pm]
		currentB := nums2[pn]
		if currentA >= currentB {
			nums1[p] = currentA
			pm--
		} else {
			nums1[p] = currentB
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
	return nums1
}

func four(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		currentM := nums1[pm]
		currentN := nums2[pn]
		if currentM >= currentN {
			nums1[p] = currentM
			pm--
		} else {
			nums1[p] = currentN
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
}

func five(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		currentM := nums1[pm]
		currentN := nums2[pn]
		if currentM >= currentN {
			nums1[p] = currentM
			pm--
		} else {
			nums1[p] = currentN
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
}

func six(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		currentM := nums1[pm]
		currentN := nums2[pn]
		if currentM >= currentN {
			nums1[p] = currentM
			pm--
		} else {
			nums1[p] = currentN
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
}

func seven(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		currentM := nums1[pm]
		currentN := nums2[pn]
		if currentM >= currentN {
			nums1[p] = currentM
			pm--
		} else {
			nums1[p] = currentN
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
}

func eight(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		if nums1[pm] > nums2[pn] {
			nums1[p] = nums1[pm]
			pm--
		} else {
			nums1[p] = nums2[pn]
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
}

func nine(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		if nums1[pm] > nums2[pn] {
			nums1[p] = nums1[pm]
			pm--
		} else {
			nums1[p] = nums2[pn]
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
}

func ten(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		if nums1[pm] > nums2[pn] {
			nums1[p] = nums1[pm]
			pm--
		} else {
			nums1[p] = nums2[pn]
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
}

func eleven(nums1 []int, nums2 []int, m, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		if nums1[pm] > nums2[pn] {
			nums1[p] = nums1[pm]
			pm--
		} else {
			nums1[p] = nums2[pn]
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
}

func twelve(nums1 []int, nums2 []int, m, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		if nums1[pm] > nums2[pn] {
			nums1[p] = nums1[pm]
			pm--
		} else {
			nums1[p] = nums2[pn]
			pn--
		}
	}
	for ; pn >= 0; pn-- {
		nums1[pn] = nums2[pn]
	}
}

func thirteen(nums1 []int, nums2 []int, m, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		if nums1[pm] > nums2[pn] {
			nums1[p] = nums1[pm]
			pm--
		} else {
			nums1[p] = nums2[pn]
			pn--
		}
	}
	for i := pn; i >= 0; i-- {
		nums1[i] = nums2[i]
	}
}

func fourteen(nums1 []int, nums2 []int, m, n int) {
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for ; pm >= 0 && pn >= 0; p-- {
		if nums1[pm] > nums2[pn] {
			nums1[p] = nums1[pm]
			pm--
		} else {
			nums1[p] = nums2[pn]
			pn--
		}
		p--
	}
	for i := pn; i >= 0; i-- {
		nums1[i] = nums2[i]

	}
}

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	m, n := 3, 3
	merge(a, m, b, n)
	fmt.Println(a)
}
