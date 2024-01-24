package main

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

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large

	for head != nil {
		current := head
		if current.Val >= x {
			largeHead.Next = current
			largeHead = largeHead.Next
		} else {
			smallHead.Next = current
			smallHead = smallHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil

	return small.Next
}

func test(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large

	for head != nil {
		current := head
		if current.Val >= x {
			largeHead.Next = current
			largeHead = largeHead.Next
		} else {
			smallHead.Next = current
			smallHead = smallHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil
	return small.Next
}

func get(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large

	for head != nil {
		current := head
		if current.Val >= x {
			largeHead.Next = current
			largeHead = largeHead.Next
		} else {
			smallHead.Next = current
			smallHead = smallHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil
	return small.Next
}

func four(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large
	for head != nil {
		if head.Val >= x {
			largeHead.Next = head
			largeHead = largeHead.Next
		} else {
			smallHead.Next = head
			smallHead = smallHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil
	return small.Next
}

func five(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	large := &ListNode{}
	smallHead := small
	largeHead := large
	for head != nil {
		if head.Val < x {
			smallHead.Next = head
			smallHead = smallHead.Next
		} else {
			largeHead.Next = head
			largeHead = largeHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil
	return small.Next
}

func six(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	large := &ListNode{}
	smallHead := small
	largeHead := large
	for head != nil {
		if head.Val < x {
			smallHead.Next = head
			smallHead = smallHead.Next
		} else {
			largeHead.Next = head
			largeHead = largeHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil
	return small.Next
}

func seven(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	large := &ListNode{}
	smallHead := small
	largeHead := large
	for head != nil {
		if head.Val < x {
			smallHead.Next = head
			smallHead = smallHead.Next
		} else {
			largeHead.Next = head
			largeHead = largeHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil
	return small.Next
}

func eight(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	large := &ListNode{}
	smallHead := small
	largeHead := large
	for head != nil {
		if head.Val < x {
			smallHead.Next = head
			smallHead = smallHead.Next
		} else {
			largeHead.Next = head
			largeHead = largeHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil
	return small.Next
}

func nine(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallHead, largeHead := small, large

	for head != nil {
		if head.Val < x {
			smallHead.Next = head
			smallHead = smallHead.Next
		} else {
			largeHead.Next = head
			largeHead = largeHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil
	return small.Next
}

func ten(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallHead, largeHead := small, large

	for head != nil {
		if head.Val < x {
			smallHead.Next = head
			smallHead = smallHead.Next
		} else {
			largeHead.Next = head
			largeHead = largeHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil
	return small.Next
}

func eleven(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallHead, largeHead := small, large

	for head != nil {
		if head.Val < x {
			smallHead.Next = head
			smallHead = smallHead.Next
		} else {
			largeHead.Next = head
			largeHead = largeHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil
	return small.Next
}

func twelve(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallHead, largeHead := small, large

	for head != nil {
		if head.Val < x {
			smallHead.Next = head
			smallHead = smallHead.Next
		} else {
			largeHead.Next = head
			largeHead = largeHead.Next
		}
		head = head.Next
	}
	smallHead.Next = large.Next
	largeHead.Next = nil

	return small.Next
}
