package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + dp[j-1]*dp[i-j]
		}
	}

	return dp[n]
}

func test(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}

func get(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}

func four(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func six(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func five(n int) int {
	if n == 0 {
		return 0
	}
	var handler func(int, int) []*TreeNode
	handler = func(start int, end int) []*TreeNode {
		var trees []*TreeNode
		if start > end {
			return append(trees, nil)
		}
		for i := start; i <= end; i++ {
			left := handler(start, i-1)
			right := handler(i+1, end)
			for _, l := range left {
				for _, r := range right {
					tree := &TreeNode{i, l, r}
					trees = append(trees, tree)
				}
			}
		}
		return trees
	}
	return len(handler(1, n))
}

func seven(n int) int {
	if n == 0 {
		return 0
	}
	var handler func(int, int) []*TreeNode
	handler = func(start int, end int) []*TreeNode {
		var trees []*TreeNode
		if start > end {
			return append(trees, nil)
		}
		for i := start; i <= end; i++ {
			left := handler(start, i-1)
			right := handler(i+1, end)
			for _, l := range left {
				for _, r := range right {
					tree := &TreeNode{i, l, r}
					trees = append(trees, tree)
				}
			}
		}
		return trees
	}
	return len(handler(1, n))

}

func nine(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func ten(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func eleven(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func twelve(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func thirteen(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func fourteen(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
}
