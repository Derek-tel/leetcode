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

func four(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	four(root.Left)
	four(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func five(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	five(root.Left)
	five(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func six(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	six(root.Left)
	six(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func seven(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	seven(root.Left)
	seven(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func eight(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	eight(root.Left)
	eight(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func nine(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	nine(root.Left)
	nine(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func ten(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	ten(root.Left)
	ten(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func eleven(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	eleven(root.Left)
	eleven(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func twelve(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	twelve(root.Left)
	twelve(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func thirteen(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	thirteen(root.Left)
	thirteen(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}
