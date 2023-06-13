package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	tail := &ListNode{}

	carry := 0

	for l1 != nil || l2 != nil {
		m, n := 0, 0
		if l1 != nil {
			m = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n = l2.Val
			l2 = l2.Next
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
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func test(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	tail := &ListNode{}

	carry := 0

	for l1 != nil || l2 != nil {
		m, n := 0, 0
		if l1 != nil {
			m = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n = l2.Val
			l2 = l2.Next
		}

		sum := m + n + carry
		carry = sum % 10
		mod := sum / 10

		if head == nil {
			head = &ListNode{Val: mod}
			tail = head
		} else {
			tail.Next = &ListNode{Val: mod}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func get(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	tail := &ListNode{}
	carry := 0

	for l1 != nil || l2 != nil {
		m, n := 0, 0
		if l1 != nil {
			m = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n = l2.Val
			l2 = l2.Next
		}
		count := m + n + carry
		carry = count / 10
		mod := count % 10
		if head == nil {
			head = &ListNode{
				Val: mod,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: mod,
			}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}
	return
}

func four(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		m, n := 0, 0
		if l1 != nil {
			m = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n = l2.Val
			l2 = l2.Next
		}
		count := m + n + carry
		carry = count / 10
		mod := count % 10
		if head == nil {
			head = &ListNode{
				Val: mod,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: mod,
			}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func five(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0

	for l1 != nil || l2 != nil {
		m, n := 0, 0
		if l1 != nil {
			m = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n = l2.Val
			l2 = l2.Next
		}
		count := m + n + carry
		carry = count / 10
		mod := count % 10
		if head == nil {
			head = &ListNode{
				Val: mod,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: mod,
			}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}
	return head
}

func six(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		m, n := 0, 0
		if l1 != nil {
			m = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n = l2.Val
			l2 = l2.Next
		}
		count := m + n + carry
		carry = count / 10
		mod := count % 10
		if head == nil {
			head = &ListNode{Val: mod}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: mod,
			}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}
	return head
}
