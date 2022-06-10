package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length <= 1 {
		return head
	}
	midNode := getMiddleNode(head)

	cur = midNode.Next
	midNode.Next = nil
	midNode = cur
	left := sortList(head)
	right := sortList(midNode)
	return mergeTwoList(left, right)
}

func getMiddleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = mergeTwoList(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoList(list1, list2.Next)
		return list2
	}
}

func test(head *ListNode) *ListNode {
	length := 0
	cur := head

	for cur != nil {
		cur = cur.Next
		length++
	}
	if length <= 1 {
		return head
	}
	mid := getMid(head)

	cur = mid.Next
	mid.Next = nil
	mid = cur
	list1 := sortList(head)
	list2 := sortList(mid)
	return merge(list1, list2)

}

func merge(listA, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	} else if listB == nil {
		return listA
	} else if listA.Val < listB.Val {
		listA.Next = merge(listA.Next, listB)
		return listA
	} else {
		listB.Next = merge(listA, listB.Next)
		return listB
	}
}

func getMid(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func get(head *ListNode) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	if length <= 1 {
		return head
	}
	mid := helper(head)
	cur = mid.Next
	mid.Next = nil
	mid = cur
	return merge(sortList(head), sortList(mid))
}

func helper(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	a := &ListNode{1, nil}
	b := &ListNode{3, a}
	c := &ListNode{2, b}
	d := &ListNode{0, c}
	sortList(d)
}
