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

func five(head *ListNode, left int, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	preHead := &ListNode{Next: head}
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

func six(head *ListNode, left int, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	preHead := &ListNode{Next: head}
	pre := preHead
	for i := 1; pre.Next != nil && i < left; i++ {
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

func seven(head *ListNode, left int, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	preHead := &ListNode{Next: head}
	pre := preHead
	for i := 1; pre.Next != nil && i < left; i++ {
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

func eight(head *ListNode, left, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	preHead := &ListNode{Next: head}
	pre := preHead
	for i := 1; pre.Next != nil && i < left; i++ {
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

func nine(head *ListNode, left, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	preHead := &ListNode{Next: head}
	pre := preHead
	for i := 1; i < left && pre.Next != nil; i++ {
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

func ten(head *ListNode, left, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}

	preHead := &ListNode{Next: head}
	prev := preHead
	for i := 1; i < left && prev.Next != nil; i++ {
		prev = prev.Next
	}
	if prev.Next == nil {
		return head
	}
	cur := prev.Next
	for i := 0; i < right-left && cur.Next != nil; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return preHead.Next
}

func eleven(head *ListNode, left, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	preHead := &ListNode{Next: head}
	prev := preHead
	for i := 1; i < left && prev.Next != nil; i++ {
		prev = prev.Next
	}
	if prev.Next == nil {
		return head
	}
	cur := prev.Next
	for i := 0; i < right-left && cur.Next != nil; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return preHead.Next
}

func twelve(head *ListNode, left, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	preHead := &ListNode{Next: head}
	prev := preHead
	for i := 1; i < left && prev.Next != nil; i++ {
		prev = prev.Next
	}
	if prev.Next == nil {
		return head
	}
	cur := prev.Next
	for i := 0; i < right-left && cur.Next != nil; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return preHead.Next
}

func thirteen(head *ListNode, left, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	preHead := &ListNode{Next: head}
	prev := preHead
	for i := 1; i < left && prev.Next != nil; i++ {
		prev = prev.Next
	}
	if prev.Next == nil {
		return head
	}
	cur := prev.Next
	for i := 0; i < right-left && cur.Next != nil; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return preHead.Next
}

func fourteen(head *ListNode, left, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	preHead := &ListNode{Next: head}
	prev := preHead
	for i := 1; i < left && prev.Next != nil; i++ {
		prev = prev.Next
	}
	if prev.Next == nil {
		return head
	}
	cur := prev.Next
	for i := 0; i < right-left && cur.Next != nil; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return preHead.Next
}

func fifteen(head *ListNode, left, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	preHead := &ListNode{Next: head}
	prev := preHead
	for i := 1; i < left && prev.Next != nil; i++ {
		prev = prev.Next
	}
	if prev.Next == nil {
		return head
	}
	cur := prev.Next
	for i := 0; i < right-left && cur.Next != nil; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return preHead.Next
}
