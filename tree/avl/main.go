package main

import "fmt"

// AVLTreeNode avl节点
type AVLTreeNode struct {
	Value  int64        // 值
	Height int64        // 该节点作为树根节点，树的高度，方便计算平衡因子
	Left   *AVLTreeNode // 左子树
	Right  *AVLTreeNode // 右字树
}

func (t *AVLTreeNode) GetHigh() int64 {
	if t == nil {
		return -1
	}
	return t.Height
}

func Max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

func (t *AVLTreeNode) LeftRotation() *AVLTreeNode {
	headNode := t.Right
	t.Right = headNode.Left
	headNode.Left = t
	//update high
	t.Height = Max(t.Left.GetHigh(), t.Right.GetHigh()) + 1
	headNode.Height = Max(headNode.Left.GetHigh(), headNode.Right.GetHigh()) + 1
	return headNode
}

func (t *AVLTreeNode) RightRotation() *AVLTreeNode {
	headNode := t.Left
	t.Left = headNode.Right
	headNode.Right = t
	//update high
	t.Height = Max(t.Left.GetHigh(), t.Right.GetHigh()) + 1
	headNode.Height = Max(headNode.Left.GetHigh(), headNode.Right.GetHigh()) + 1
	return headNode
}

func (t *AVLTreeNode) RightLeftRotation() *AVLTreeNode {
	sonNode := t.Right.RightRotation()
	t.Right = sonNode
	return t.LeftRotation()
}

func (t *AVLTreeNode) LeftRightRotation() *AVLTreeNode {
	sonNode := t.Left.LeftRotation()
	t.Left = sonNode
	return t.RightRotation()
}

func (t *AVLTreeNode) Adjust() *AVLTreeNode {
	if t.Right.GetHigh()-t.Left.GetHigh() == 2 {
		if t.Right.Right.GetHigh() > t.Right.Left.GetHigh() {
			t.LeftRotation()
		} else {
			t.RightLeftRotation()
		}
	} else if t.Left.GetHigh()-t.Right.GetHigh() == 2 {
		if t.Left.Left.GetHigh() > t.Left.Right.GetHigh() {
			t.RightRotation()
		} else {
			t.LeftRightRotation()
		}
	}
	return t
}

func (t *AVLTreeNode) Insert(value int64) *AVLTreeNode {
	if t == nil {
		return &AVLTreeNode{value, 1, nil, nil}
	}
	if value < t.Value {
		t.Left = t.Left.Insert(value)
		t = t.Adjust()
	} else if value > t.Value {
		t.Right = t.Right.Insert(value)
		t = t.Adjust()
	} else {
		fmt.Println("this node exist")
	}
	t.Height = Max(t.Left.GetHigh(), t.Right.GetHigh()) + 1
	return t
}

func (t *AVLTreeNode) Delete(value int64) *AVLTreeNode {
	if t == nil {
		return t
	}
	compare := value - t.Value
	if compare < 0 {
		t.Left = t.Left.Delete(value)
	} else if compare > 0 {
		t.Right = t.Right.Delete(value)
	} else {
		if t.Right != nil && t.Left != nil {
			t.Value = t.Right.GetMin()
			t.Right = t.Right.Delete(t.Value)
		} else if t.Left != nil {
			t = t.Left
		} else {
			t = t.Right
		}
	}
	if t != nil {
		t.Height = Max(t.Left.GetHigh(), t.Right.GetHigh()) + 1
		t.Adjust()
	}
	return t
}

func (t *AVLTreeNode) GetMin() int64 {
	if t == nil {
		return -1
	}
	if t.Left == nil {
		return t.Value
	}
	return t.Left.GetMin()
}
