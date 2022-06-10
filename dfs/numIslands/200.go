package numIslands

func numIslands(grid [][]byte) int {

	if len(grid) == 0 {
		return 0
	}
	length := len(grid)
	weight := len(grid[0])
	count := 0

	for i := 0; i < length; i++ {
		for j := 0; j < weight; j++ {
			if grid[i][j]-'0' == 1 {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, x, y int) {
	length := len(grid)
	weight := len(grid[0])

	grid[x][y] = '0'
	if x-1 >= 0 && grid[x-1][y] == '1' {
		dfs(grid, x-1, y)
	}
	if x+1 < length && grid[x+1][y] == '1' {
		dfs(grid, x+1, y)
	}
	if y-1 >= 0 && grid[x][y-1] == '1' {
		dfs(grid, x, y-1)
	}
	if y+1 < weight && grid[x][y+1] == '1' {
		dfs(grid, x, y+1)
	}
}

func test(grid [][]byte) int {
	h, w := len(grid), len(grid[0])
	count := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j]-'0' == 1 {
				count++
				dfs1(grid, i, j)
			}
		}
	}
	return count
}

func dfs1(grid [][]byte, x, y int) {
	grid[x][y] = 0
	h, w := len(grid), len(grid[0])

	if x-1 >= 0 && grid[x-1][y] == 1 {
		dfs(grid, x-1, y)
	}
	if x+1 < h && grid[x+1][y] == 1 {
		dfs(grid, x+1, y)
	}
	if y-1 >= 0 && grid[x][y-1] == 1 {
		dfs(grid, x, y-1)
	}
	if y+1 < w && grid[x][y+1] == 1 {
		dfs(grid, x, y+1)
	}

}
