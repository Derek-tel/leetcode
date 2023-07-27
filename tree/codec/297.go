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

type Codec struct {
	input         []string
	stringBuilder strings.Builder
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.stringBuilder.WriteString("null,")
		return ""
	}
	this.stringBuilder.WriteString(strconv.Itoa(root.Val) + ",")
	this.serialize(root.Left)
	this.serialize(root.Right)
	return this.stringBuilder.String()
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
	node := &TreeNode{
		Val:   value,
		Left:  this.helper(),
		Right: this.helper(),
	}
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
