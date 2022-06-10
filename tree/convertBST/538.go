package convertBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(*TreeNode)
	sum := 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum = sum + node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func test(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum = sum + node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func get(root *TreeNode) *TreeNode {
	sum := 0
	var helper func(*TreeNode)

	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Right)
		sum = sum + node.Val
		node.Val = sum
		helper(node.Left)
	}
	helper(root)
	return root
}
