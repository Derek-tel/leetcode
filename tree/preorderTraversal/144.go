package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	list := []int{}
	preOrder(root, &list)
	return list
}

func preOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	preOrder(root.Left, list)
	preOrder(root.Right, list)
}

func test(root *TreeNode) []int {
	list := []int{}
	helper(root, &list)
	return list
}

func helper(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	helper(root.Left, list)
	helper(root.Right, list)
}

func get(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

func four(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	for len(stack) > 0 || root != nil {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Right
	}
	return result
}

func five(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	for len(stack) > 0 || root != nil {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Right
	}
	return result
}

func six(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	for len(stack) > 0 || root != nil {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Right
	}
	return result
}

func seven(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	for len(stack) > 0 || root != nil {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Right
	}
	return result
}
