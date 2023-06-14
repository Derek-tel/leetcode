package removeNthFromEnd

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dumpy := &ListNode{0, head}
	slow := dumpy
	fast := head

	for i := n; i > 0; i-- {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dumpy.Next
}

func test(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dumpy := &ListNode{-1, head}
	slow := dumpy
	fast := head

	for i := n; i > 0; i-- {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dumpy.Next

}

func get(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	preHead := &ListNode{Val: -1, Next: head}
	slow := preHead
	fast := head
	for i := n; i > 0; i-- {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return preHead.Next
}

func four(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	preHead := &ListNode{Val: -1, Next: head}
	slow := preHead
	fast := head
	for i := n; i > 0; i-- {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return preHead.Next
}

func five(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	preHead := &ListNode{Next: head}
	slow := preHead
	fast := head
	for i := n; i > 0; i-- {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return preHead.Next
}

func six(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	preHead := &ListNode{Next: head}
	slow := preHead
	fast := head
	for i := n; i > 0; i-- {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return preHead.Next
}
