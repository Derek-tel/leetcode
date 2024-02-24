package main

func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	//init
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for i := 0; i < col; i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func demo(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for j := 0; j < col; j++ {
		if j == 0 {
			dp[0][j] = grid[0][j]
		} else {
			dp[0][j] = grid[0][j] + dp[0][j-1]
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func test(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for i := 0; i < col; i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func get(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for i := 0; i < col; i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func five(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for j := 0; j < col; j++ {
		if j == 0 {
			dp[0][j] = grid[0][j]
		} else {
			dp[0][j] = dp[0][j-1] + grid[0][j]
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func six(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[0][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for j := 0; j < col; j++ {
		if j == 0 {
			dp[0][j] = grid[0][0]
		} else {
			dp[0][j] = dp[0][j-1] + grid[0][j]
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func seven(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[0][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for j := 0; j < col; j++ {
		if j == 0 {
			dp[0][j] = grid[0][0]
		} else {
			dp[0][j] = dp[0][j-1] + grid[0][j]
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func eight(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for i := 0; i < col; i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func nine(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for i := 0; i < col; i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[row-1][col-1]
}

func ten(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for i := 0; i < col; i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func eleven(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[0][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for i := 0; i < col; i++ {
		if i == 0 {
			dp[0][i] = grid[0][0]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func twelve(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[0][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for i := 0; i < col; i++ {
		if i == 0 {
			dp[0][i] = grid[0][0]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func thirteen(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		if i == 0 {
			dp[i][0] = grid[0][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for i := 0; i < col; i++ {
		if i == 0 {
			dp[0][i] = grid[0][0]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
