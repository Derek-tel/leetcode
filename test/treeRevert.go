package main

import "fmt"

type TreeNode struct {
	Val  int
	Next *TreeNode
}

func revertTree(head *TreeNode) *TreeNode {
	if head == nil {
		return head
	}

	preHead := &TreeNode{-1, head}
	pre := preHead
	current := head
	for current.Next != nil {
		next := current.Next
		current.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return preHead.Next
}

func main() {
	node1 := &TreeNode{4, nil}
	node2 := &TreeNode{3, node1}
	node3 := &TreeNode{2, node2}
	node4 := &TreeNode{1, node3}
	res := revertTree(node4)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
