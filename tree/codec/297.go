package main

import (
	"strconv"
	"strings"
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

type Codec struct {
	Input   []string
	Builder strings.Builder
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.Builder.WriteString("null,")
		return ""
	}
	this.Builder.WriteString(strconv.Itoa(root.Val) + ",")
	this.serialize(root.Left)
	this.serialize(root.Right)
	return this.Builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	this.Input = strings.Split(data, ",")
	return this.helper()
}

func (this *Codec) helper() *TreeNode {
	if this.Input[0] == "null" {
		this.Input = this.Input[1:]
		return nil
	}
	value, _ := strconv.Atoi(this.Input[0])
	this.Input = this.Input[1:]
	root := &TreeNode{value, nil, nil}
	root.Left = this.helper()
	root.Right = this.helper()

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
