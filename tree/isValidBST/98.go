package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	min := math.Inf(-1)
	max := math.Inf(1)
	return isValid(root, min, max)
}

func isValid(root *TreeNode, min float64, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v > min && v < max && isValid(root.Left, min, v) && isValid(root.Right, v, max)
}

func isValidBST1(root *TreeNode) bool {
	list := []int{}
	inorder(root, &list)
	for i := 1; i < 1; i++ {
		if list[i-1] > list[i] {
			return false
		}
	}
	return false
}

func inorder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, list)
	*list = append(*list, root.Val)
	inorder(root.Right, list)
}

func main() {
	min := math.Inf(-1)
	max := math.Inf(1)
	fmt.Println(min, max)

	for i := 1; i < 1; i++ {
		fmt.Println("xxx")
	}
}

func test(root *TreeNode) bool {
	max := math.MaxInt
	min := math.MinInt
	return check(root, min, max)
}

func check(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	return root.Val > min && root.Val < max && check(root.Left, min, root.Val) && check(root.Right, root.Val, max)
}

func get(root *TreeNode) bool {
	max := math.MaxInt
	min := math.MinInt
	return helper(root, max, min)
}

func helper(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	return root.Val > min && root.Val < max && helper(root.Left, root.Val, min) && helper(root.Right, max, root.Val)
}

func five(root *TreeNode) bool {
	var handler func(*TreeNode, float64, float64) bool
	handler = func(node *TreeNode, min float64, max float64) bool {
		if node == nil {
			return true
		}
		v := float64(node.Val)
		return v > min && v < max && handler(node.Left, min, v) && handler(node.Right, v, max)
	}
	return handler(root, math.Inf(-1), math.Inf(1))
}

func six(root *TreeNode) bool {
	var handler func(*TreeNode, int, int) bool
	handler = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		return node.Val > min && node.Val < max && handler(node.Left, min, node.Val) && handler(node.Right, node.Val, max)
	}
	return handler(root, math.MinInt, math.MaxInt)
}

func seven(root *TreeNode) bool {
	var handler func(*TreeNode, int, int) bool
	handler = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		return node.Val > min && node.Val < max && handler(node.Left, min, node.Val) && handler(node.Right, node.Val, max)
	}
	return handler(root, math.MinInt, math.MaxInt)
}

func eight(root *TreeNode) bool {
	var handler func(*TreeNode, int, int) bool
	handler = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		return node.Val > min && node.Val < max && handler(node.Left, min, node.Val) && handler(node.Right, node.Val, max)
	}
	return handler(root, math.MinInt, math.MaxInt)
}

func nine(root *TreeNode) bool {
	var handler func(*TreeNode, int, int) bool
	handler = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		return node.Val > min && node.Val < max && handler(node.Left, min, node.Val) && handler(node.Right, node.Val, max)
	}
	return handler(root, math.MinInt, math.MaxInt)
}

func ten(root *TreeNode) bool {
	var handler func(*TreeNode, int, int) bool
	handler = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		return node.Val > min && node.Val < max && handler(node.Left, min, node.Val) && handler(node.Right, node.Val, max)
	}
	return handler(root, math.MinInt, math.MaxInt)
}

func eleven(root *TreeNode) bool {
	var handler func(*TreeNode, int, int) bool
	handler = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		return node.Val > min && node.Val < max && handler(node.Left, min, node.Val) && handler(node.Right, node.Val, max)
	}
	return handler(root, math.MinInt, math.MaxInt)
}

func twelve(root *TreeNode) bool {
	var handler func(*TreeNode, int, int) bool
	handler = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		return node.Val > min && node.Val < max && handler(node.Left, min, node.Val) && handler(node.Right, node.Val, max)
	}
	return handler(root, math.MinInt, math.MaxInt)
}

func thirteen(root *TreeNode) bool {
	var handler func(*TreeNode, int, int) bool
	handler = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		return node.Val > min && node.Val < max && handler(node.Left, min, node.Val) && handler(node.Right, node.Val, max)
	}
	return handler(root, math.MinInt, math.MaxInt)
}

func fourteen(root *TreeNode) bool {
	var handler func(*TreeNode, int, int) bool
	handler = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		return node.Val > min && node.Val < max && handler(node.Left, min, node.Val) && handler(node.Right, node.Val, max)
	}
	return handler(root, math.MinInt, math.MaxInt)
}
