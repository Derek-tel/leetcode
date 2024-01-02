package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	input   []string
	builder strings.Builder
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.builder.WriteString("null,")
		return ""
	}
	this.builder.WriteString(strconv.Itoa(root.Val) + ",")
	this.serialize(root.Left)
	this.serialize(root.Right)
	return this.builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	this.input = strings.Split(data, ",")
	return this.helper()
}

func (this *Codec) helper() *TreeNode {
	if this.input[0] == "null" {
		this.input = this.input[1:]
		return nil
	}
	value, _ := strconv.Atoi(this.input[0])
	this.input = this.input[1:]
	return &TreeNode{
		Val:   value,
		Left:  this.helper(),
		Right: this.helper()}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
