package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func test(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(test(root.Left), test(root.Right)) + 1
}

func get(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(get(root.Left), get(root.Right)) + 1
}
