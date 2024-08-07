package lowestCommonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func test(root, q, p *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := test(root.Left, q, p)
	right := test(root.Right, q, p)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func get(root, q, p *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := get(root.Left, q, p)
	right := get(root.Right, q, p)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func four(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := four(root.Left, p, q)
	right := four(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func five(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := five(root.Left, p, q)
	right := five(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func six(root, q, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := six(root.Left, p, q)
	right := six(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func seven(root, q, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := seven(root.Left, q, p)
	right := seven(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func eight(root, q, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := eight(root.Left, q, p)
	right := eight(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func nine(root, q, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := nine(root.Left, q, p)
	right := nine(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func ten(root, q, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := ten(root.Left, q, p)
	right := ten(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func eleven(root, q, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := eleven(root.Left, q, p)
	right := eleven(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return root
	}
	return left
}

func twelve(root, q, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := twelve(root.Left, q, p)
	right := twelve(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func thirteen(root, q, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := thirteen(root.Left, q, p)
	right := thirteen(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func fourteen(root, q, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := fourteen(root.Left, q, p)
	right := fourteen(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
