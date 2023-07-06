package postorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	var result []int
	var prev *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Right == nil || top.Right == prev {
			result = append(result, top.Val)
			prev = top
		} else {
			stack = append(stack, top)
			root = top.Right
		}
	}
	return result
}

func one(root *TreeNode) []int {
	stack := []*TreeNode{}
	var result []int
	var prev *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Right == nil || top.Right == prev {
			result = append(result, top.Val)
			prev = top
		} else {
			stack = append(stack, top)
			root = top.Right
		}
	}
	return result
}

func two(root *TreeNode) []int {
	stack := []*TreeNode{}
	var result []int
	var prev *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Right == nil || top.Right == prev {
			result = append(result, top.Val)
			prev = top
		} else {
			stack = append(stack, top)
			root = top.Right
		}
	}
	return result
}

func three(root *TreeNode) []int {
	stack := []*TreeNode{}
	var result []int
	var prev *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Right == nil || top.Right == prev {
			result = append(result, top.Val)
			prev = top
		} else {
			stack = append(stack, top)
			root = top.Right
		}
	}
	return result
}
