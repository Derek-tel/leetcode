package main

import "fmt"

type Node2 struct {
	Value int
	Next  *Node2
}

func mergeTwoList(listA, listB *Node2) *Node2 {
	preNode := &Node2{}
	pre := preNode
	A := listA
	B := listB
	for A != nil && B != nil {
		if A.Value < B.Value {
			pre.Next = A
			A = A.Next
		} else {
			pre.Next = B
			B = B.Next
		}
		pre = pre.Next
	}
	if A == nil {
		pre.Next = B
	} else {
		pre.Next = A
	}
	return preNode.Next
}

func main() {
	node1 := &Node2{2, nil}
	node2 := &Node2{1, node1}
	node3 := &Node2{5, nil}
	node4 := &Node2{4, node3}
	node5 := &Node2{3, node4}

	res := mergeTwoList(node2, node5)
	for res != nil {
		fmt.Println(res.Value)
		res = res.Next
	}
	fmt.Println("------")
	node1 = &Node2{2, nil}
	node2 = &Node2{1, node1}
	res1 := mergeTwoList(node2, nil)
	for res1 != nil {
		fmt.Println(res1.Value)
		res1 = res1.Next
	}
}
