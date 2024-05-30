package rob

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob337(root *TreeNode) int {
	if root == nil {
		return 0
	}
	noRob, doRob := dfs(root)
	if noRob > doRob {
		return noRob
	}
	return doRob
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftNoRob, leftRob := dfs(root.Left)
	rightNoRob, rightRob := dfs(root.Right)

	//当前节点打劫,左右节点不能被打劫，左右相加
	doRob := root.Val + leftNoRob + rightNoRob
	//当前节点不打劫，左右节点 可以打劫，也可以不打，返回的价值是左右相加
	noRob := max(leftRob, leftNoRob) + max(rightRob, rightNoRob)
	return noRob, doRob
}

func demo(root *TreeNode) int {
	if root == nil {
		return 0
	}
	noRob, doRob := help(root)
	if noRob > doRob {
		return noRob
	}
	return doRob
}

func help(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftNoRob, leftDoRob := help(root.Left)
	rightNoRob, rightDoRob := help(root.Right)

	doRob := root.Val + leftNoRob + rightNoRob
	noRob := max(leftDoRob, leftNoRob) + max(rightDoRob, rightNoRob)
	return noRob, doRob
}

func test1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	rob, noRob := helper(root)

	return max(rob, noRob)
}

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	robLeft, noRobLeft := helper(root.Left)
	robRight, noRobRight := helper(root.Right)

	//当前节点打劫  左右不能打劫
	rob := root.Val + noRobLeft + noRobRight

	//当前节点不打劫，左右节点可以打，也可以不打

	noRob := max(robLeft, noRobLeft) + max(robRight, noRobRight)

	return rob, noRob

}

func get(root *TreeNode) int {
	if root == nil {
		return 0
	}
	rob, noRob := getter(root)
	return max(rob, noRob)
}

func getter(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	robLeft, noRobLeft := getter(root.Left)
	robRight, noRobRight := getter(root.Right)

	rob := root.Val + noRobLeft + noRobRight
	noRob := max(robLeft, noRobLeft) + max(robRight, noRobRight)

	return rob, noRob
}

func five(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		robLeft, notRobLeft := handler(node.Left)
		robRight, notRobRight := handler(node.Right)

		do := node.Val + notRobRight + notRobLeft
		notDo := max(robLeft, notRobLeft) + max(robRight, notRobRight)

		return do, notDo
	}
	doRob, notRob := handler(root)
	return max(doRob, notRob)
}

func six(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		robLeft, notRobLeft := handler(node.Left)
		robRight, notRobRight := handler(node.Right)
		do := node.Val + notRobRight + notRobLeft
		notDo := max(robLeft, notRobLeft) + max(robRight, notRobRight)

		return do, notDo
	}
	doRob, doNotRob := handler(root)
	return max(doRob, doNotRob)
}

func seven(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		robLeft, notRobLeft := handler(node.Left)
		robRight, notRobRight := handler(node.Right)

		doRob := node.Val + notRobRight + notRobLeft
		doNotRob := max(robLeft, notRobLeft) + max(robRight, notRobRight)

		return doRob, doNotRob
	}
	doRob, doNotRob := handler(root)
	return max(doRob, doNotRob)
}

func eight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		robLeft, notRobLeft := handler(node.Left)
		robRight, norRobRight := handler(node.Right)

		doRob := node.Val + norRobRight + notRobLeft
		notRob := max(robLeft, notRobLeft) + max(robRight, norRobRight)

		return doRob, notRob
	}
	do, doNot := handler(root)
	return max(do, doNot)
}

func nine(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		robLeft, notRobLeft := handler(node.Left)
		robRight, notRobRight := handler(node.Right)

		doRob := node.Val + notRobRight + notRobLeft
		notRob := max(robLeft, notRobLeft) + max(robRight, notRobRight)

		return doRob, notRob
	}

	doRob, notRob := handler(root)
	return max(doRob, notRob)
}

func ten(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		robLeft, notRobLeft := handler(node.Left)
		robRight, notRobRight := handler(node.Right)

		doRob := node.Val + notRobRight + notRobLeft
		notRob := max(robLeft, notRobLeft) + max(robRight, notRobRight)

		return doRob, notRob
	}
	doRob, notRob := handler(root)
	return max(doRob, notRob)
}

func twelve(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		doRobLeft, notRobLeft := handler(node.Left)
		doRobRight, notRobRight := handler(node.Right)

		doRob := node.Val + notRobRight + notRobLeft
		notRob := max(doRobLeft, notRobLeft) + max(doRobRight, notRobRight)

		return doRob, notRob
	}
	doRob, notRob := handler(root)
	return max(doRob, notRob)
}
func thirteen(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		doRobLeft, notRobLeft := handler(node.Left)
		doRobRight, notRobRight := handler(node.Right)

		doRob := node.Val + notRobRight + notRobLeft
		notRob := max(doRobLeft, notRobLeft) + max(doRobRight, notRobRight)
		return doRob, notRob
	}
	doRob, notRob := handler(root)
	return max(doRob, notRob)
}

func fourteen(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		doRobLeft, notRobLeft := handler(node.Left)
		doRobRight, notRobRight := handler(node.Right)

		doRob := node.Val + notRobLeft + notRobRight
		notRob := max(doRobLeft, notRobLeft) + max(doRobRight, notRobRight)

		return doRob, notRob
	}

	doRob, notRob := handler(root)
	return max(doRob, notRob)
}

func fifteen(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		doRobLeft, notRobLeft := handler(node.Left)
		doRobRight, notRobRight := handler(node.Right)

		doRob := node.Val + notRobRight + notRobLeft
		notRob := max(doRobLeft, notRobLeft) + max(doRobRight, notRobRight)

		return doRob, notRob
	}
	doRob, notRob := handler(root)
	return max(doRob, notRob)
}

func sixteen(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		doRobLeft, notRobLeft := handler(node.Left)
		doRobRight, notRobRight := handler(node.Right)

		doRob := node.Val + notRobRight + notRobLeft
		notRob := max(doRobLeft, notRobLeft) + max(doRobRight, notRobRight)

		return doRob, notRob
	}

	doRob, notRob := handler(root)
	return max(doRob, notRob)
}

func seventeen(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var handler func(*TreeNode) (int, int)
	handler = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		doRobLeft, notRobLeft := handler(node.Left)
		doRobRight, notRobRight := handler(node.Right)

		doRob := node.Val + notRobRight + notRobLeft
		notRob := max(doRobLeft, notRobLeft) + max(doRobRight, notRobRight)

		return doRob, notRob
	}
	doRob, notRob := handler(root)
	return max(doRob, notRob)
}
