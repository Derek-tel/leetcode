package main

type ListItem struct {
	Val  int
	Next *ListItem
}

func removeDuplicate(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dumpy := sortList(head)
	cur := dumpy
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dumpy
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var getMid func(*ListNode) *ListNode
	getMid = func(node *ListNode) *ListNode {
		fast, slow := node, node
		for fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next.Next
		}
		return slow
	}
	count := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		count++
	}
	if count <= 1 {
		return head
	} else {
		mid := getMid(head)
		next := mid.Next
		mid.Next = nil
		return mergeTwo(sortList(head), sortList(next))
	}
}

func mergeTwo(left, right *ListNode) *ListNode {
	preHead := &ListNode{}
	prev := preHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			prev.Next = left
			left = left.Next
		} else {
			prev.Next = right
			right = right.Next
		}
		prev = prev.Next
	}
	if left != nil {
		prev.Next = left
	} else {
		prev.Next = right
	}
	return preHead.Next
}
