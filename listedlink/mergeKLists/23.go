package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return mergeTwoList(lists[0], lists[1])
	} else {
		mid := len(lists) / 2
		return mergeTwoList(mergeKLists(lists[0:mid]), mergeKLists(lists[mid:]))
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

func test(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return merge(lists[0], lists[1])
	} else {
		mid := len(lists) / 2
		return merge(test(lists[:mid]), test(lists[mid:]))
	}
}

func merge(listA *ListNode, listB *ListNode) *ListNode {
	prehead := &ListNode{Val: -1}
	prev := prehead
	for listA != nil && listB != nil {
		if listA.Val <= listB.Val {
			prev.Next = listA
			listA = listA.Next
		} else {
			prev.Next = listB
			listB = listB.Next
		}
		prev = prev.Next
	}
	if listA == nil {
		prev.Next = listB
	} else {
		prev.Next = listA
	}
	return prehead.Next
}

func get(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return help(lists[0], lists[1])
	} else {
		mid := len(lists) / 2
		return help(get(lists[0:mid]), get(lists[mid:]))
	}
}

func helper(listA, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	} else if listB == nil {
		return listA
	} else if listA.Val < listB.Val {
		listA.Next = helper(listA.Next, listB)
		return listA
	} else {
		listB.Next = helper(listA, listB.Next)
		return listB
	}
}

func help(listA, listB *ListNode) *ListNode {
	preHead := &ListNode{Val: -1}
	prev := preHead
	for listA != nil && listB != nil {
		if listA.Val < listB.Val {
			prev.Next = listA
			listA = listA.Next
		} else {
			prev.Next = listB
			listB = listB.Next
		}
		prev = prev.Next
	}
	if listA == nil {
		prev.Next = listB
	}
	if listB == nil {
		prev.Next = listA
	}
	return preHead.Next
}

func four(lists []*ListNode) *ListNode {
	var handler func(*ListNode, *ListNode) *ListNode
	handler = func(listA *ListNode, listB *ListNode) *ListNode {
		preHead := &ListNode{Val: -1}
		pre := preHead
		for listA != nil && listB != nil {
			if listA.Val < listB.Val {
				pre.Next = listA
				listA = listA.Next
			} else {
				pre.Next = listB
				listB = listB.Next
			}
			pre = pre.Next
		}
		if listA == nil {
			pre.Next = listB
		} else {
			pre.Next = listA
		}
		return preHead.Next
	}
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return handler(lists[0], lists[1])
	} else {
		mid := len(lists) / 2
		return handler(four(lists[0:mid]), four(lists[mid:]))
	}
}
