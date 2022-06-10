package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	list := preorderTraversal(root)
	cur := &TreeNode{}
	cur = root
	for i := 1; i < len(list); i++ {
		cur.Left = nil
		cur.Right = &TreeNode{list[i], nil, nil}
		cur = cur.Right
	}
	return
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

func test(root *TreeNode) {
	curr := root

	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}

func get(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr.Right
			curr.Left = nil
			curr.Right = next
		}
		curr = curr.Right
	}
}
