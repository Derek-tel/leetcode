package main

import (
	"fmt"
	sort2 "sort"
)

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	length := len(intervals)
	if length == 0 {
		return res
	}
	sort2.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	temp := intervals[0]
	for i := 1; i < length; i++ {
		if temp[1] < intervals[i][0] { //分割的
			res = append(res, temp)
			temp = intervals[i]
			continue
		} else {
			if temp[1] >= intervals[i][1] {
				continue
			} else {
				temp[1] = intervals[i][1]
			}
		}
	}
	res = append(res, temp)
	return res
}
func four(intervals [][]int) [][]int {
	var result [][]int
	sort2.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if temp[1] < intervals[i][0] {
			result = append(result, temp)
			temp = intervals[i]
		} else {
			if temp[1] >= intervals[i][1] {
				continue
			} else {
				temp[1] = intervals[i][1]
			}
		}
	}
	result = append(result, temp)
	return result
}

func test(intervals [][]int) [][]int {
	res := [][]int{}
	sort(intervals, 0, len(intervals)-1)
	temp := intervals[0]

	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > temp[1] {
			res = append(res, temp)
			temp = intervals[i]
		} else {
			if temp[1] > intervals[i][1] {
				continue
			} else {
				temp[1] = intervals[i][1]
			}
		}
	}
	res = append(res, temp)
	return res
}

func sort(intervals [][]int, start, end int) {
	temp := intervals[start]
	p := start

	i, j := start, end

	for i <= j {
		if j >= p && intervals[j][0] >= temp[0] {
			j--
		}
		if j >= p {
			intervals[j], intervals[p] = intervals[p], intervals[j]
			p = j
		}

		if i <= p && intervals[i][0] < temp[0] {
			i++
		}
		if i <= p {
			intervals[i], intervals[p] = intervals[p], intervals[i]
			p = i
		}
	}
	intervals[p] = temp
	if p-start > 1 {
		sort(intervals, start, p-1)
	}
	if end-p > 1 {
		sort(intervals, p+1, end)
	}
}

func quickSort(intervals [][]int, start, end int) {
	temp := intervals[start]
	p := start
	i, j := start, end

	for i < j {
		if i < j && intervals[j][0] >= temp[0] {
			j--
		}
		if j >= p {
			intervals[j], intervals[p] = intervals[p], intervals[j]
			p = j
		}

		if i < j && intervals[i][0] < temp[0] {
			i++
		}
		if i < p {
			intervals[i], intervals[p] = intervals[p], intervals[i]
			p = i
		}
	}
	intervals[p] = temp
	if p-start > 1 {
		quickSort(intervals, start, p-1)
	}
	if end-p > 1 {
		quickSort(intervals, p+1, end)
	}

}

func main() {
	test := [][]int{
		{2, 6},
		{1, 3},
		{8, 10},
		{7, 11},
	}
	fmt.Println(merge(test))
}
