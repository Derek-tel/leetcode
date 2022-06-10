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
