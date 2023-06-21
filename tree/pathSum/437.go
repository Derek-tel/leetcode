package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	if root.Val == targetSum {
		res = +1
	}
	res += rootSum(root.Left, targetSum-root.Val)
	res += rootSum(root.Right, targetSum-root.Val)
	return
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func test(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	res := helper(root, target)
	res += test(root.Left, target)
	res += test(root.Right, target)
	return res
}

func helper(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	res := 0

	if root.Val == target {
		res += 1
	}
	res += helper(root.Left, target-root.Val)
	res += helper(root.Right, target-root.Val)
	return res
}

func get(root *TreeNode, target int) int {
	res := help(root, target)
	res += get(root.Left, target)
	res += get(root.Right, target)
	return res
}

func help(root *TreeNode, target int) int {
	res := 0
	if root == nil {
		return res
	}
	if root.Val == target {
		res += 1
	}
	res += help(root.Left, target-root.Val)
	res += help(root.Right, target-root.Val)
	return res
}

func four(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode, int) int
	handler = func(node *TreeNode, target int) int {
		result := 0
		if node == nil {
			return 0
		}
		if node.Val == target {
			result += 1
		}
		result += handler(node.Left, target-node.Val)
		result += handler(node.Right, target-node.Val)
		return result
	}
	resp := 0
	resp += handler(root, targetSum)
	resp += four(root.Left, targetSum)
	resp += four(root.Right, targetSum)
	return resp
}

func five(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode, int) int
	handler = func(node *TreeNode, target int) int {
		result := 0
		if node == nil {
			return 0
		}
		if node.Val == target {
			result += 1
		}
		result += handler(node.Left, target-node.Val)
		result += handler(node.Right, target-node.Val)
		return result
	}
	resp := 0
	resp += handler(root, targetSum)
	resp += five(root.Left, targetSum)
	resp += five(root.Right, targetSum)
	return resp
}

func six(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode, int) int
	handler = func(node *TreeNode, target int) int {
		resp := 0
		if node == nil {
			return resp
		}
		if node.Val == target {
			resp += 1
		}
		resp += handler(node.Left, target-node.Val)
		resp += handler(node.Right, target-node.Val)
		return resp
	}
	result := 0
	result += handler(root, targetSum)
	result += six(root.Left, targetSum)
	result += six(root.Right, targetSum)
	return result
}
