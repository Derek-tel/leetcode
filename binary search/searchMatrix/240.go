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

func second(matrix [][]int, target int) bool {
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

func three(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func four(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func five(matrix [][]int, target int) bool {
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

func six(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func seven(matrix [][]int, target int) bool {
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

func eight(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func nine(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func ten(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func eleven(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func twelve(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func thirteen(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func fourteen(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func fifteen(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func sixteen(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func seventeen(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}
