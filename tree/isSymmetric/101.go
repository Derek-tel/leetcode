package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSame(invert(root.Left), root.Right)
}

func isSame(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		return p.Val == q.Val && isSame(p.Left, q.Left) && isSame(p.Right, q.Right)
	} else {
		return false
	}
}

func invert(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invert(root.Left)
	invert(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func test(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && helper(left.Left, right.Right) && helper(left.Right, right.Left)
}

func get(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func charge(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right != nil {
		return left.Val == right.Val && charge(left.Left, right.Right) && charge(left.Right, right.Left)
	} else {
		return false
	}
}

func five(root *TreeNode) bool {
	var handler func(*TreeNode, *TreeNode) bool
	handler = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && handler(left.Left, right.Right) && handler(left.Right, right.Left)
	}
	return handler(root.Left, root.Right)
}

func six(root *TreeNode) bool {
	var handler func(*TreeNode, *TreeNode) bool
	handler = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && handler(left.Left, right.Right) && handler(left.Right, right.Left)
	}

	return handler(root.Left, root.Right)
}

func seven(root *TreeNode) bool {
	var handler func(*TreeNode, *TreeNode) bool
	handler = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && handler(left.Left, right.Right) && handler(left.Right, right.Left)
	}

	return handler(root.Left, root.Right)
}

func eight(root *TreeNode) bool {
	var handler func(*TreeNode, *TreeNode) bool
	handler = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && handler(left.Left, right.Right) && handler(left.Right, right.Left)
	}
	return handler(root.Left, root.Right)
}
