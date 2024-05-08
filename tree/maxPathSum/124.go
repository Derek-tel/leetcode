package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	getMax(root, &max)
	return max
}

func getMax(root *TreeNode, max *int) int {
	if root == nil {
		return math.MinInt32
	}

	left := getMax(root.Left, max)
	right := getMax(root.Right, max)

	maxCurr := Max(Max(left+root.Val, right+root.Val), root.Val)

	*max = Max(*max, Max(left+right+root.Val, maxCurr))
	return maxCurr
}

func test(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	helper(root, &max)
	return max
}

func helper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, max)
	right := helper(root.Right, max)

	currMax := Max(Max(left+root.Val, right+root.Val), root.Val)
	*max = Max(*max, currMax)
	return currMax
}

func get(root *TreeNode) int {
	max := math.MinInt32
	getIt(root, &max)
	return max
}

func getIt(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := getIt(root.Left, max)
	right := getIt(root.Right, max)
	currMax := Max(Max(left+root.Val, right+root.Val), root.Val)
	*max = Max(*max, Max(left+right+root.Val, currMax))
	return currMax
}

func four(root *TreeNode) int {
	max := math.MinInt32
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(max, Max(left+right+node.Val, currMax))
		return currMax
	}
	handler(root)
	return max
}

func five(root *TreeNode) int {
	max := math.MinInt
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(max, Max(left+right+node.Val, currMax))
		return currMax
	}
	handler(root)
	return max
}

func six(root *TreeNode) int {
	max := math.MinInt
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(Max(currMax, max), left+right+node.Val)
		return currMax
	}
	handler(root)
	return max
}

func seven(root *TreeNode) int {
	max := math.MinInt32
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(max, Max(currMax, left+right+node.Val))
		return currMax
	}
	handler(root)
	return max
}

func eight(root *TreeNode) int {
	max := math.MinInt
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(max, Max(currMax, left+right+node.Val))
		return currMax
	}
	handler(root)
	return max
}

func nine(root *TreeNode) int {
	max := math.MinInt
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(max, Max(currMax, left+right+node.Val))
		return currMax
	}
	handler(root)
	return max
}

func ten(root *TreeNode) int {
	max := math.MinInt
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(max, Max(currMax, left+right+node.Val))
		return currMax
	}
	handler(root)
	return max
}

func eleven(root *TreeNode) int {
	max := math.MinInt
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(max, Max(currMax, left+right+node.Val))
		return currMax
	}
	handler(root)
	return max
}

func twelve(root *TreeNode) int {
	max := math.MinInt32
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(max, Max(currMax, left+node.Val+right))
		return currMax
	}
	handler(root)
	return max
}

func thirteen(root *TreeNode) int {
	max := math.MinInt
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currentMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(max, Max(currentMax, left+node.Val+right))
		return currentMax
	}
	handler(root)
	return max
}

func fourteen(root *TreeNode) int {
	max := math.MaxInt
	var handler func(*TreeNode) int
	handler = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := handler(node.Left)
		right := handler(node.Right)
		currentMax := Max(Max(left+node.Val, right+node.Val), node.Val)
		max = Max(max, Max(currentMax, left+node.Val+right))
		return currentMax
	}
	handler(root)
	return max
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
