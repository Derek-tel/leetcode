package rbtree

import (
	"github.com/welllog/data-structures/utils"
)

type color bool

const (
	black, red color = true, false
)

type Node struct {
	key    interface{}
	value  interface{}
	color  color
	left   *Node
	right  *Node
	parent *Node
}
type Tree struct {
	root       *Node
	size       int
	comparator utils.Comparator
}

func NewWith(comparator utils.Comparator) *Tree {
	return &Tree{comparator: comparator}
}

func NewWithIntComparator() *Tree {
	return &Tree{comparator: utils.IntComparator}
}

func NewWithStringComparator() *Tree {
	return &Tree{comparator: utils.StringComparator}
}

func (t *Tree) Put(key, value interface{}) {
	var insertNode *Node
	if t.root == nil {
		t.comparator(key, key)
		t.root = &Node{key: key, value: value, color: red}
		insertNode = t.root
	} else {
		node := t.root
		loop := true
		for loop {
			c := t.comparator(key, node.key)
			switch {
			case c == 0:
				node.key = key
				node.value = value
				return
			case c < 0:
				if node.left == nil {
					node.left = &Node{key: key, value: value, color: red, parent: node}
					insertNode = node.left
					loop = false
				} else {
					node = node.left
				}
			case c > 0:
				if node.right == nil {
					node.right = &Node{key: key, value: value, color: red, parent: node}
					insertNode = node.right
					loop = false
				} else {
					node = node.right
				}
			}
		}
		insertNode.parent = node
	}
	t.insertCase1(insertNode)
}

func (t *Tree) insertCase1(node *Node) {
	if node.parent == nil {
		node.color = black
		return
	}
	t.insertCase2(node)
}

func (t *Tree) insertCase2(node *Node) {
	if node.parent.color == black {
		return
	}
	t.insertCase3(node)
}

func (t *Tree) insertCase3(node *Node) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.parent.color = black
		uncle.color = black
		grandparent := node.grandparent()
		grandparent.color = red
		t.insertCase1(grandparent)
	} else {
		t.insertCase4(node)
	}
}

func (t *Tree) insertCase4(node *Node) {
	grandParent := node.grandparent()
	if node == node.parent.right && node.parent == grandParent.left {
		t.rotateLeft(node.parent)
		node = node.left
	} else if node == node.parent.left && node.parent == grandParent.right {
		t.rotationRight(node.parent)
		node = node.right
	}
	t.insertCase5(node)
}

func (t *Tree) insertCase5(node *Node) {
	node.parent.color = black
	grandParent := node.grandparent()
	grandParent.color = red
	if node == node.parent.left && node.parent == node.parent.left {
		t.rotationRight(grandParent)
	} else if node == node.parent.right && node.parent == node.parent.right {
		t.rotateLeft(grandParent)
	}
}

func (t *Tree) remove(key interface{}) {
	var child *Node
	node := t.lookup(key)
	if node == nil {
		return
	}
	if node.left != nil && node.right != nil {
		pred := node.left.maximumNode()
		node.key = pred.key
		node.value = pred.value
		node = pred
	}
	if node.left == nil {
		child = node.right
	} else {
		child = node.left
	}
	if node.color == black {
		if nodeColor(child) == red {
			child.color = black
		} else {
			t.deleteCase1(node)
		}
	}
	t.replaceNode(node, child)
	t.size--
}

func (t *Tree) deleteCase1(node *Node) {
	if node.parent == nil {
		return
	}
	t.deleteCase2(node)
}

func (t *Tree) deleteCase2(node *Node) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.parent.color = red
		sibling.color = black
		if node == node.parent.left {
			t.rotateLeft(node.parent)
		} else {
			t.rotationRight(node.parent)
		}
	}
	t.deleteCase3(node)
}

// 父节点，兄弟节点，兄弟节点的子节点都为黑色。将兄弟节点绘为红色,在父节点上重新平衡
func (t *Tree) deleteCase3(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.left) == black &&
		nodeColor(sibling.right) == black {
		sibling.color = red
		t.deleteCase1(node.parent)
	} else {
		t.deleteCase4(node)
	}
}

// 父节点为红色，兄弟节点跟兄弟节点的子节点都为黑色，将父节点置为黑色，兄弟节点置为红色,重新达到平衡
func (t *Tree) deleteCase4(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.left) == black &&
		nodeColor(sibling.right) == black {
		sibling.color = red
		node.parent.color = black
	} else {
		t.deleteCase5(node)
	}
}

// 兄弟节点为黑色。当前节点为左孩子，兄弟节点左孩子为红色；当前节点为右孩子，兄弟节点右孩子为红色，兄弟节点进行右旋或左旋，转为下一种情况
func (t *Tree) deleteCase5(node *Node) {
	sibling := node.sibling()
	// 节点为左孩子，兄弟节点为黑色，兄弟节点的左孩子为红色，在兄弟节点上做右旋
	// 兄弟节点的左孩子成为了兄弟节点的父节点，当前节点的新兄弟。交换旧的兄弟节点跟新兄弟节点的颜色
	if node == node.parent.left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.left) == red &&
		nodeColor(sibling.right) == black {
		sibling.color = red
		sibling.left.color = black
		t.rotationRight(sibling)
	} else if node == node.parent.right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.right) == red &&
		nodeColor(sibling.left) == black {
		sibling.color = red
		sibling.right.color = black
		t.rotateLeft(sibling)
	}
	// 进入下一种处理
	t.deleteCase6(node)
}

// 兄弟节点为黑色。当前节点为左孩子，兄弟节点的右孩子为红色；当前节点为右孩子，兄弟节点左孩子为红色.在父节点上左旋或右旋
func (t *Tree) deleteCase6(node *Node) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.parent)
	node.parent.color = black
	if node == node.parent.left && nodeColor(sibling.right) == red {
		sibling.right.color = black
		t.rotateLeft(node.parent)
	} else if nodeColor(sibling.left) == red {
		sibling.left.color = black
		t.rotationRight(node.parent)
	}
}

func (t *Tree) replaceNode(old *Node, new *Node) {
	if old.parent == nil {
		t.root = new
	} else {
		if old == old.parent.left {
			old.parent.left = new
		} else {
			old.parent.right = new
		}
	}
	if new != nil {
		new.parent = old.parent
	}
}

func (n *Node) maximumNode() *Node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (t *Tree) lookup(key interface{}) *Node {
	node := t.root
	for node != nil {
		c := t.comparator(key, node.key)
		switch {
		case c == 0:
			return node
		case c > 0:
			node = node.right
		case c < 0:
			node = node.left
		}
	}
	return nil
}

func (t *Tree) rotateLeft(node *Node) {
	right := node.right
	if node.parent != nil {
		if node == node.parent.left {
			node.parent.left = right
		} else {
			node.parent.right = right
		}
	} else {
		t.root = right
	}
	right.parent = node.parent
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	right.left = node
	node.parent = right
}

func (t *Tree) rotationRight(node *Node) {
	left := node.left
	if node.parent != nil {
		if node == node.parent.left {
			node.parent.left = left
		} else {
			node.parent.right = left
		}
	} else {
		t.root = left
	}
	left.parent = node.parent
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	left.right = node
	node.parent = left
}

func (n *Node) grandparent() *Node {
	if n != nil && n.parent != nil {
		return n.parent.parent
	}
	return nil
}

func (n *Node) uncle() *Node {
	if n == nil || n.parent == nil || n.parent.parent == nil {
		return nil
	}
	return n.parent.sibling()
}

func (n *Node) sibling() *Node {
	if n == nil || n.parent == nil {
		return nil
	}
	if n == n.parent.left {
		return n.parent.right
	}
	return n.parent.left
}

func nodeColor(node *Node) color {
	if node == nil {
		return black
	}
	return node.color
}
