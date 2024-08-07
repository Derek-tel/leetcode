package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	if head == nil {
		return true
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	half := slow
	secondHalf := reverse(half.Next)

	p1 := head
	p2 := secondHalf

	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func test(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	half := slow
	secondHalf := reverse1(half.Next)

	p1 := head
	p2 := secondHalf

	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func reverse1(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func get(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	half := slow
	secondHalf := helper(half.Next)
	p1 := head
	p2 := secondHalf
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func helper(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func four(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast, slow := head, head
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		var pre *ListNode
		cur := node
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		return pre
	}
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	half := slow
	second := handler(half.Next)
	p1 := head
	p2 := second
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func five(head *ListNode) bool {
	if head == nil {
		return true
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		var pre *ListNode
		cur := node
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		return pre
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	p1 := head
	p2 := handler(slow.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func six(head *ListNode) bool {
	if head == nil {
		return true
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		var pre *ListNode
		cur := node
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		return pre
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1 := head
	p2 := handler(slow.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func seven(head *ListNode) bool {
	if head == nil {
		return true
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}
		var dumpy = &ListNode{Next: head}
		pre := dumpy
		cur := node
		for cur.Next != nil {
			next := cur.Next
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}
		return dumpy.Next
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	p1 := head
	p2 := handler(slow.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func eight(head *ListNode) bool {
	if head == nil {
		return false
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		if node == nil {
			return nil
		}
		var dumpy = &ListNode{Next: node}
		pre := dumpy
		cur := node
		for cur.Next != nil {
			next := cur.Next
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}
		return dumpy.Next
	}
	fast, slot := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slot = slot.Next
	}
	p1 := head
	p2 := handler(slot.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func nine(head *ListNode) bool {
	if head == nil {
		return false
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}
		var dumpy = &ListNode{Next: node}
		prev := dumpy
		cur := node
		for cur.Next != nil {
			next := cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		return dumpy.Next
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1 := head
	p2 := handler(slow.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func ten(head *ListNode) bool {
	if head == nil {
		return false
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}
		var dumpy = &ListNode{Next: node}
		prev := dumpy
		cur := node
		for cur.Next != nil {
			next := cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		return dumpy.Next
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1 := head
	p2 := handler(slow.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func eleven(head *ListNode) bool {
	if head == nil {
		return false
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}
		dumpy := &ListNode{Next: node}
		prev := dumpy
		cur := node
		for cur.Next != nil {
			next := cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		return dumpy.Next
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1 := head
	p2 := handler(slow.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func twelve(head *ListNode) bool {
	if head == nil {
		return false
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}
		dumpy := &ListNode{Next: node}
		prev := dumpy
		cur := node
		for cur.Next != nil {
			next := cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		return dumpy.Next
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1 := head
	p2 := handler(slow.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func thirteen(head *ListNode) bool {
	if head == nil {
		return false
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}
		dumpy := &ListNode{Next: node}
		prev := dumpy
		cur := node
		for cur.Next != nil {
			next := cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		return dumpy.Next
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1 := head
	p2 := handler(slow.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func fourteen(head *ListNode) bool {
	if head == nil {
		return false
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}
		dumpy := &ListNode{Next: node}
		prev := dumpy
		cur := node
		for cur.Next != nil {
			next := cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		return dumpy.Next
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1 := head
	p2 := handler(slow.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func fifteen(head *ListNode) bool {
	if head == nil {
		return false
	}
	var handler func(*ListNode) *ListNode
	handler = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}
		dumpy := &ListNode{Next: node}
		prev := dumpy
		cur := node
		for cur.Next != nil {
			next := cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		return dumpy.Next
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1 := head
	p2 := handler(slow.Next)
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}
