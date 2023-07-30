package recoverTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pre *TreeNode

	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && top.Val < pre.Val {
			y = top
			if x == nil {
				x = pre
			} else {
				break
			}
		}
		pre = top
		root = top.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

func test(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pre *TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && top.Val < pre.Val {
			y = top
			if x == nil {
				x = pre
			} else {
				break
			}
		}
		pre = top
		root = top.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

func get(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	var x, y, pre *TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && top.Val < pre.Val {
			y = top
			if x == nil {
				x = pre
			} else {
				break
			}
		}
		pre = top
		root = top.Right
	}
	x.Val, y.Val = y.Val, x.Val
	return root
}

func four(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	var x, y, prev *TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && top.Val < prev.Val {
			y = top
			if x == nil {
				x = prev
			} else {
				break
			}
		}
		root = top.Right
		prev = top
	}
	x.Val, y.Val = y.Val, x.Val
	return root
}

func five(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	var x, y, pre *TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && top.Val < pre.Val {
			y = top
			if x == nil {
				x = pre
			} else {
				break
			}
		}
		root = top.Right
		pre = top
	}
	x.Val, y.Val = y.Val, x.Val
	return root
}

func six(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	var x, y, pre *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && top.Val < pre.Val {
			y = top
			if x == nil {
				x = pre
			} else {
				break
			}
		}
		root = top.Right
		pre = top
	}
	x.Val, y.Val = y.Val, x.Val
	return root
}

func seven(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	var x, y, pre *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && top.Val < pre.Val {
			y = top
			if x == nil {
				x = pre
			} else {
				break
			}
		}
		root = top.Right
		pre = top
	}
	x.Val, y.Val = y.Val, x.Val
	return root
}

func eight(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	var x, y, prev *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && top.Val < prev.Val {
			y = top
			if x == nil {
				x = prev
			} else {
				break
			}
		}
		prev = top
		root = top.Right
	}
	x.Val, y.Val = y.Val, x.Val
	return root
}
