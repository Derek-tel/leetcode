package diameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode) int
	res := 0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		res = max(res, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return res
}

func test(root *TreeNode) int {
	sum := 0
	var helper func(*TreeNode) int

	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)
		sum = max(sum, left+right)
		return max(left, right) + 1
	}
	helper(root)
	return sum
}

func get(root *TreeNode) int {
	sum := 0
	var helper func(*TreeNode) int

	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)
		sum = max(sum, left+right)
		return max(left, right) + 1
	}
	helper(root)
	return sum
}

func four(root *TreeNode) int {
	sum := 0
	var helper func(*TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)
		sum = max(sum, left+right)
		return max(left, right) + 1
	}
	helper(root)
	return sum
}

func five(root *TreeNode) int {
	sum := 0
	var helper func(*TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)
		sum = max(sum, left+right)
		return max(left, right) + 1
	}
	helper(root)
	return sum
}

func six(root *TreeNode) int {
	sum := 0
	var helper func(*TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)
		sum = max(sum, left+right)
		return max(right, left) + 1
	}
	helper(root)
	return sum
}

func seven(root *TreeNode) int {
	sum := 0
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		sum = max(sum, left+right)
		return max(left, right) + 1
	}
	handler(root)
	return sum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
