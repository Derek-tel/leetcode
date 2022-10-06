package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left > right {
		return head
	}

	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	cur := &ListNode{}

	for i := 0; pre.Next != nil && i < left-1; i++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}

	cur = pre.Next
	for i := 0; i < right-left; i++ {
		nex := cur.Next
		cur.Next = nex.Next
		nex.Next = pre.Next
		pre.Next = nex
	}
	return newHead.Next
}

func test(head *ListNode, left int, right int) *ListNode {
	if head == nil || left < right {
		return nil
	}
	newHead := &ListNode{-1, head}
	pre := newHead
	for i := 0; pre.Next != nil && i < left-1; i++ {
		pre = pre.Next
	}

	if pre.Next == nil {
		return head
	}

	cur := pre.Next

	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return newHead.Next
}

func get(head *ListNode, left, right int) *ListNode {
	if head == nil || left <= right {
		return head
	}
	preHead := &ListNode{-1, head}
	pre := preHead
	for i := 0; pre.Next != nil && i < left-1; i++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next

	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return preHead.Next
}

func four(head *ListNode, left int, right int) *ListNode {
	if head == nil || left > right {
		return head
	}
	preHead := &ListNode{-1, head}
	pre := preHead
	for i := 0; pre.Next != nil && i < left-1; i++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return preHead.Next
}
