package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var inOrder = func(node *TreeNode) {}

	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		res = append(res, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	return res
}

func inorderTraversal2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		root = top.Right
	}
	return
}

func test(root *TreeNode) []int {
	res := []int{}
	var stack []*TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		root = top.Right
	}
	return res
}

func get(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		root = top.Right
	}
	return res
}

func five(root *TreeNode) []int {
	var inorder func(*TreeNode)
	var result []int
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result
}

func six(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.Val)
		root = top.Right
	}
	return result
}

func seven(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.Val)
		root = top.Right
	}
	return result
}

func eight(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.Val)
		root = top.Right
	}
	return result
}

func nine(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.Val)
		root = top.Right
	}
	return result
}
func main() {
	b := &TreeNode{3, nil, nil}
	c := &TreeNode{2, b, nil}
	a := &TreeNode{1, nil, c}
	fmt.Println(inorderTraversal2(a))
}
