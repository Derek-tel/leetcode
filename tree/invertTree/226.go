package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	invert(root)
	return root
}

func invert(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invert(root.Left)
	invert(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func test(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	test(root.Left)
	test(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}

func get(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	helper(root.Left)
	helper(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
