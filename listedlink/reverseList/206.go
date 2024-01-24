package reverseList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dumpy := &ListNode{0, head}
	pre := dumpy
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dumpy.Next
}

func test(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dumpy := &ListNode{0, head}
	pre := dumpy
	cur := head

	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dumpy.Next
}

func three(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	preHead := &ListNode{-1, head}
	pre := preHead
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return preHead.Next
}

func four(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumpy := &ListNode{Next: head}
	pre := dumpy
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dumpy.Next
}

func five(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dumpy := &ListNode{Next: head}
	prev := dumpy
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dumpy.Next
}

func six(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumpy := &ListNode{Next: head}
	prev := dumpy
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dumpy.Next
}

func seven(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumpy := &ListNode{Next: head}
	prev := dumpy
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dumpy.Next
}

func eight(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumpy := &ListNode{Next: head}
	prev := dumpy
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dumpy.Next
}

func nine(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumpy := &ListNode{Next: head}
	prev := dumpy
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dumpy.Next
}

func ten(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumpy := &ListNode{Next: head}
	prev := dumpy
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dumpy.Next
}

func eleven(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumpy := &ListNode{Next: head}
	prev := dumpy
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dumpy.Next
}
