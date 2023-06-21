package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return merge(root1, root2)
}

func merge(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root1.Val = root1.Val + root2.Val
	root1.Left = merge(root1.Left, root2.Left)
	root1.Right = merge(root1.Right, root2.Right)
	return root1
}

func test(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val = root1.Val + root2.Val
	root1.Left = test(root1.Left, root2.Left)
	root1.Right = test(root1.Right, root2.Right)
	return root1
}

func get(rootQ, rootP *TreeNode) *TreeNode {
	if rootQ == nil {
		return rootP
	}
	if rootP == nil {
		return rootQ
	}

	rootQ.Val = rootP.Val + rootQ.Val
	rootQ.Left = get(rootQ.Left, rootP.Left)
	rootQ.Right = get(rootQ.Right, rootP.Right)
	return rootQ
}

func four(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val + root2.Val
	root1.Left = four(root1.Left, root2.Left)
	root1.Right = four(root1.Right, root2.Right)
	return root1
}

func five(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val + root2.Val
	root1.Left = five(root1.Left, root2.Left)
	root1.Right = five(root1.Right, root2.Right)
	return root1
}

func six(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val + root2.Val
	root1.Left = six(root1.Left, root2.Left)
	root1.Right = six(root1.Right, root2.Right)
	return root1
}
