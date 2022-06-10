package main

func searchMatrix(matrix [][]int, target int) bool {
	length := len(matrix)
	if length == 0 {
		return false
	}

	for _, col := range matrix {
		low, high := 0, len(col)-1
		for low <= high {
			mid := low + (high-low)>>1
			if col[mid] > target {
				high = mid - 1
			} else if col[mid] < target {
				low = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}
