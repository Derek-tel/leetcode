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

func five(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(five(root.Left), five(root.Right)) + 1
}

func six(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(six(root.Left), six(root.Right)) + 1
}

func seven(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := seven(root.Left)
	right := seven(root.Right)
	return max(left, right) + 1
}

func eight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := eight(root.Left)
	right := eight(root.Right)
	return max(left, right) + 1
}

func nine(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := nine(root.Left)
	right := nine(root.Right)
	return max(left, right) + 1
}

func ten(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := ten(root.Left)
	right := ten(root.Right)
	return max(left, right) + 1
}

func eleven(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := ten(root.Left)
	right := ten(root.Right)
	return max(left, right) + 1
}

func twelve(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := twelve(root.Left)
	right := twelve(root.Right)
	return max(left, right) + 1
}

func thirteen(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := thirteen(root.Left)
	right := thirteen(root.Right)
	return max(left, right) + 1
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
