package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoList(listA, listB *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for listA != nil || listB != nil {
		m, n := 0, 0
		if listA != nil {
			m = listA.Val
			listA = listA.Next
		}
		if listB != nil {
			n = listB.Val
			listB = listB.Next
		}

		sum := m + n + carry
		mod := sum % 10
		carry = sum / 10
		if head == nil {
			head = &ListNode{Val: mod}
			tail = head
		} else {
			tail.Next = &ListNode{Val: mod}
			tail = tail.Next
		}
	}

	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func reverse(list *ListNode) *ListNode {
	return list
}
