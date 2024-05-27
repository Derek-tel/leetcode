package generateMatrix

func one(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	row, col, dirIndex := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := pairList[dirIndex]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIndex = (dirIndex + 1) % 4
			dir = pairList[dirIndex]
		}
		row = row + dir.x
		col = col + dir.y
	}
	return matrix
}

func two(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}
	row, col, dirIndex := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := pairList[dirIndex]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIndex = (dirIndex + 1) % 4
			dir = pairList[dirIndex]
		}
		row = row + dir.x
		col = col + dir.y
	}
	return matrix
}

func three(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}
	row, col, dirIndex := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := pairList[dirIndex]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIndex = (dirIndex + 1) % 4
			dir = pairList[dirIndex]
		}
		row = row + dir.x
		col = col + dir.y
	}
	return matrix
}

func four(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	row, col, dirIndex := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := pairList[dirIndex]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIndex = (dirIndex + 1) % 4
			dir = pairList[dirIndex]
		}
		row = row + dir.x
		col = col + dir.y
	}
	return matrix
}

func five(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	row, col, dirIndex := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := pairList[dirIndex]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIndex = (dirIndex + 1) % 4
			dir = pairList[dirIndex]
		}
		row = row + dir.x
		col = col + dir.y
	}
	return matrix
}

type pair struct {
	x, y int
}

var pairList = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
