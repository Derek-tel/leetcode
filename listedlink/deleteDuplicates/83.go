package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func one(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func two(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func three(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func removeDuplicate(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	find := make(map[int]bool)
	find[cur.Val] = true
	for cur.Next != nil {
		if _, ok := find[cur.Next.Val]; ok {
			cur.Next = cur.Next.Next
		} else {
			find[cur.Next.Val] = true
			cur = cur.Next
		}
	}
	return head
}

func four(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {
	node4 := &ListNode{1, nil}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{3, node3}
	node1 := &ListNode{4, node2}

	head := removeDuplicate(node1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
