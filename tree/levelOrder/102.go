package main

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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		l := len(queue)
		temp := []int{}
		for i := 0; i < l; i++ {
			top := queue[i]
			temp = append(temp, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		queue = queue[l:]
		res = append(res, temp)
	}
	return res
}

func levelOrderDfs(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	var getOrder = func(root *TreeNode, level int) {}
	getOrder = func(root *TreeNode, level int) {
		if len(res) > level {
			res[level] = append(res[level], root.Val)
		} else {
			res = append(res, []int{root.Val})
		}

		if root.Left != nil {
			getOrder(root.Left, level+1)
		}
		if root.Right != nil {
			getOrder(root.Right, level+1)
		}
	}
	getOrder(root, 0)
	return res
}

func test(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		l := len(queue)
		temp := []int{}
		for i := 0; i < l; i++ {
			top := queue[i]
			temp = append(temp, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		queue = queue[l:]
		res = append(res, temp)
	}
	return res
}

func get(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		temp := []int{}
		for i := 0; i < length; i++ {
			top := queue[i]
			temp = append(temp, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		queue = queue[length:]
		res = append(res, temp)
	}
	return res
}

func dfs(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) > level {
			res[level] = append(res[level], node.Val)
		} else {
			res = append(res, []int{node.Val})
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return res
}

func six(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			top := queue[i]
			temp = append(temp, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}
func seven(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			tag := queue[i]
			temp = append(temp, tag.Val)
			if tag.Left != nil {
				queue = append(queue, tag.Left)
			}
			if tag.Right != nil {
				queue = append(queue, tag.Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}

func eight(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			tag := queue[i]
			temp = append(temp, tag.Val)
			if tag.Left != nil {
				queue = append(queue, tag.Left)
			}
			if tag.Right != nil {
				queue = append(queue, tag.Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}

func ten(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			tag := queue[i]
			temp = append(temp, tag.Val)
			if tag.Left != nil {
				queue = append(queue, tag.Left)
			}
			if tag.Right != nil {
				queue = append(queue, tag.Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}

func eleven(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}

func twelve(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}

func thirteen(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}

func fourteen(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}

func fifteen(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}

func sixteen(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}

func seventeen(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, temp)
		queue = queue[length:]
	}
	return result
}
