package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{rootVal, nil, nil}
	for i, v := range inorder {
		if v == rootVal {
			root.Left = buildTree(preorder[1:i+1], inorder[0:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
		}
	}
	return root
}

func test(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{rootVal, nil, nil}
	for i, v := range inorder {
		if v == rootVal {
			root.Left = test(preorder[1:i+1], inorder[0:i])
			root.Right = test(preorder[i+1:], inorder[i:])
		}
	}
	return root
}

func get(preOrder []int, inorder []int) *TreeNode {
	if len(preOrder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preOrder[0]
	root := &TreeNode{rootVal, nil, nil}
	for i, v := range inorder {
		if v == rootVal {
			root.Left = get(preOrder[1:i+1], inorder[0:i])
			root.Right = get(preOrder[i+1:], inorder[i+1:])
		}
	}
	return root
}

func four(preOrder []int, inorder []int) *TreeNode {
	if len(preOrder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preOrder[0]
	root := &TreeNode{rootVal, nil, nil}
	for i, v := range inorder {
		if v == rootVal {
			root.Left = four(preOrder[1:i+1], inorder[0:i])
			root.Right = four(preOrder[i+1:], inorder[i+1:])
		}
	}
	return root
}

func five(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{rootVal, nil, nil}
	for i, v := range inorder {
		if v == rootVal {
			root.Left = five(preorder[1:i+1], inorder[0:i])
			root.Right = five(preorder[i+1:], inorder[i+1:])
		}
	}
	return root
}

func six(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{rootVal, nil, nil}
	for i, v := range inorder {
		if v == rootVal {
			root.Left = six(preorder[1:i+1], inorder[0:i])
			root.Right = six(preorder[i+1:], inorder[i+1:])
		}
	}
	return root
}

func seven(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{rootVal, nil, nil}
	for i, v := range inorder {
		if v == rootVal {
			root.Left = seven(preorder[1:i+1], inorder[0:i])
			root.Right = seven(preorder[i+1:], inorder[i+1:])
		}
	}
	return root
}
