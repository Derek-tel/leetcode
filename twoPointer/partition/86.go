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
