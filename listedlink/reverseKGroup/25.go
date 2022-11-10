package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	flag := tail.Next
	p := head
	for flag != tail {
		tmp := p.Next
		p.Next = flag
		flag = p
		p = tmp
	}
	return tail, head
}

func test(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	dumpy := &ListNode{Next: head}
	pre := dumpy

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dumpy.Next
			}
		}
		next := tail.Next
		head, tail = reverse1(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return dumpy.Next
}

func reverse1(head, tail *ListNode) (*ListNode, *ListNode) {
	dumpy := &ListNode{0, head}
	pre := dumpy
	cur := head
	flag := tail.Next
	for cur.Next != flag {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dumpy.Next, head

	/*flag := tail.Next
	pre := head
	for flag != tail {
		next := pre.Next
		pre.Next = flag
		flag = pre
		pre = next
	}
	return tail, head*/
}

func get(head *ListNode, k int) *ListNode {
	preHead := &ListNode{Next: head}
	pre := preHead

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return preHead.Next
			}
		}
		next := tail.Next
		head, tail = helper(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return preHead.Next
}

func helper(head, tail *ListNode) (*ListNode, *ListNode) {
	preHead := &ListNode{Next: head}
	pre := preHead
	cur := head
	flag := tail.Next
	for cur.Next != flag {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return preHead.Next, head
}

func four(head *ListNode, k int) *ListNode {
	preHead := &ListNode{Next: head}
	pre := preHead
	var handler func(*ListNode, *ListNode) (*ListNode, *ListNode)
	handler = func(first *ListNode, last *ListNode) (*ListNode, *ListNode) {
		dumpy := &ListNode{Next: first}
		prev := dumpy
		cur := first
		flag := last.Next
		for cur.Next != flag {
			next := cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		return dumpy.Next, first
	}
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return preHead.Next
			}
		}
		next := tail.Next
		head, tail = handler(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return preHead.Next
}

func five(head *ListNode, k int) *ListNode {
	preHead := &ListNode{Next: head}
	pre := preHead

	var handler func(*ListNode, *ListNode) (*ListNode, *ListNode)
	handler = func(first *ListNode, last *ListNode) (*ListNode, *ListNode) {
		dumpy := &ListNode{Next: first}
		prev := dumpy
		cur := first
		flag := last.Next
		for cur.Next != flag {
			next := cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		return dumpy.Next, cur
	}

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return preHead.Next
			}
		}
		next := tail.Next
		head, tail = handler(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return preHead.Next
}
