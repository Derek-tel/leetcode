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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	prevHead := &ListNode{Val: -1}
	prev := prevHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}
	return prevHead.Next
}

func five(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{Val: -1}
	prev := preHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}
	return preHead.Next
}

func test(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = test(list1.Next, list2)
		return list1
	} else {
		list2.Next = test(list1, list2.Next)
		return list2
	}
}

func four(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = four(list1.Next, list2)
		return list1
	} else {
		list2.Next = four(list1, list2.Next)
		return list2
	}
}

func six(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	pre := preHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 == nil {
		pre.Next = list2
	} else {
		pre.Next = list1
	}
	return preHead.Next
}

func seven(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	pre := preHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 == nil {
		pre.Next = list2
	} else {
		pre.Next = list1
	}
	return preHead.Next
}

func eight(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	pre := preHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	} else {
		pre.Next = list2
	}
	return preHead.Next
}

func nine(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	pre := preHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	} else {
		pre.Next = list2
	}
	return preHead.Next
}

func ten(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	pre := preHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}

	if list1 != nil {
		pre.Next = list1
	} else {
		pre.Next = list2
	}
	return preHead.Next
}

func eleven(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	pre := preHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	} else {
		pre.Next = list2
	}
	return preHead.Next
}

func twelve(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	prev := preHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}
	return preHead.Next
}

func thirteen(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	prev := preHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}
	return preHead.Next
}

func fourteen(list1, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	prev := preHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}
	return preHead.Next
}

func fifteen(list1, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	prev := preHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}
	return preHead.Next
}
