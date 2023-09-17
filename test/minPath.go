/**
- 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。每次只能向下或者向右移动一步。输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPath(matrix))
}
func minPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = matrix[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + matrix[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + matrix[i][j]
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}

	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
