package main

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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func generate(start int, end int) []*TreeNode {
	tree := []*TreeNode{}
	if start > end {
		return append(tree, nil)
	}
	for i := start; i <= end; i++ {
		left := generate(start, i-1)
		right := generate(i+1, end)
		for _, l := range left {
			for _, r := range right {
				node := &TreeNode{Val: i, Left: l, Right: r}
				tree = append(tree, node)
			}
		}
	}

	return tree
}

func test(n int) []*TreeNode {
	return gene(1, n)
}

func gene(start int, end int) []*TreeNode {
	tree := []*TreeNode{}
	if start > end {
		return append(tree, nil)
	}

	for i := start; i <= end; i++ {
		left := gene(start, i-1)
		right := gene(i+1, end)
		for _, l := range left {
			for _, r := range right {
				node := &TreeNode{i, l, r}
				tree = append(tree, node)
			}
		}
	}
	return tree
}

func get(n int) []*TreeNode {
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	tree := []*TreeNode{}
	if start > end {
		return append(tree, nil)
	}
	for i := start; i <= end; i++ {
		left := helper(start, i-1)
		right := helper(i+1, end)
		for _, l := range left {
			for _, r := range right {
				node := &TreeNode{i, l, r}
				tree = append(tree, node)
			}
		}
	}
	return tree
}

func four(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var handler func(int, int) []*TreeNode
	handler = func(start, end int) []*TreeNode {
		tree := []*TreeNode{}
		if start > end {
			return append(tree, nil)
		}
		for i := start; i <= end; i++ {
			left := handler(start, i-1)
			right := handler(i+1, end)
			for _, l := range left {
				for _, r := range right {
					node := &TreeNode{i, l, r}
					tree = append(tree, node)
				}
			}
		}
		return tree
	}

	return handler(1, n)
}

func five(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var handler func(int, int) []*TreeNode
	handler = func(start int, end int) []*TreeNode {
		tree := []*TreeNode{}
		if start > end {
			return append(tree, nil)
		}
		for i := start; i <= end; i++ {
			left := handler(start, i-1)
			right := handler(i+1, end)
			for _, l := range left {
				for _, r := range right {
					flag := &TreeNode{i, l, r}
					tree = append(tree, flag)
				}
			}
		}
		return tree
	}
	return handler(1, n)
}

func six(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var handler func(int, int) []*TreeNode
	handler = func(start int, end int) []*TreeNode {
		tree := []*TreeNode{}
		if start > end {
			return append(tree, nil)
		}
		for i := start; i <= end; i++ {
			left := handler(start, i-1)
			right := handler(i+1, end)
			for _, l := range left {
				for _, r := range right {
					flag := &TreeNode{i, l, r}
					tree = append(tree, flag)
				}
			}
		}
		return tree
	}
	return handler(1, n)
}
