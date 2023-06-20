package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	builder strings.Builder
	input   []string
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
	fmt.Println(this.builder.String())
	return this.builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	this.input = strings.Split(data, ",")
	return this.deserializeHelper()
}

func (this *Codec) deserializeHelper() *TreeNode {
	if this.input[0] == "null" {
		this.input = this.input[1:]
		return nil
	}
	value, _ := strconv.Atoi(this.input[0])
	this.input = this.input[1:]
	return &TreeNode{Val: value,
		Left:  this.deserializeHelper(),
		Right: this.deserializeHelper(),
	}
}

type code struct {
	input   []string
	builder strings.Builder
}

func construct() code {
	return code{}
}

func (c *code) encode(tree *TreeNode) string {
	if tree == nil {
		c.builder.WriteString("null,")
		return ""
	}
	c.builder.WriteString(strconv.Itoa(tree.Val) + ",")
	c.encode(tree.Left)
	c.encode(tree.Right)
	return c.builder.String()
}

func (c *code) decode(str string) *TreeNode {
	if len(str) == 0 {
		return nil
	}
	c.input = strings.Split(str, ",")
	return c.helper()
}

func (c *code) helper() *TreeNode {
	if c.input[0] == "null" {
		c.input = c.input[1:]
		return nil
	}
	val, _ := strconv.Atoi(c.input[0])
	c.input = c.input[1:]
	return &TreeNode{
		Val:   val,
		Left:  c.helper(),
		Right: c.helper(),
	}
}

type Coupe struct {
	Input   []string
	Builder strings.Builder
}

func constructor() Coupe {
	return Coupe{}
}

func (c *Coupe) encode(tree *TreeNode) string {
	if tree == nil {
		c.Builder.WriteString("null,")
		return ""
	}
	c.Builder.WriteString(strconv.Itoa(tree.Val) + ",")
	c.encode(tree.Left)
	c.encode(tree.Right)
	return c.Builder.String()
}

func (c *Coupe) decode(str string) *TreeNode {
	if len(str) == 0 {
		return nil
	}
	c.Input = strings.Split(str, ",")
	return c.helper()
}

func (c *Coupe) helper() *TreeNode {
	if c.Input[0] == "null" {
		c.Input = c.Input[1:]
		return nil
	}
	val, _ := strconv.Atoi(c.Input[0])
	c.Input = c.Input[1:]
	return &TreeNode{
		Val:   val,
		Left:  c.helper(),
		Right: c.helper(),
	}
}

type FourCouple struct {
	Input   []string
	Builder strings.Builder
}

func FourConstructor() FourCouple {
	return FourCouple{}
}

func (c *FourCouple) serialize(tree *TreeNode) string {
	if tree == nil {
		c.Builder.WriteString("null,")
		return ""
	}
	c.Builder.WriteString(strconv.Itoa(tree.Val) + ",")
	c.serialize(tree.Left)
	c.serialize(tree.Right)
	return c.Builder.String()
}

func (c *FourCouple) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	c.Input = strings.Split(data, ",")
	return c.helper()
}

func (c *FourCouple) helper() *TreeNode {
	if c.Input[0] == "null" {
		c.Input = c.Input[1:]
		return nil
	}
	val, _ := strconv.Atoi(c.Input[0])
	c.Input = c.Input[1:]
	return &TreeNode{
		Val:   val,
		Left:  c.helper(),
		Right: c.helper(),
	}
}

type FiveCouple struct {
	Input   []string
	Builder strings.Builder
}

func FiveConstructor() FiveCouple {
	return FiveCouple{}
}

func (c *FiveCouple) serialize(tree *TreeNode) string {
	if tree == nil {
		c.Builder.WriteString("null,")
		return ""
	}
	c.Builder.WriteString(strconv.Itoa(tree.Val) + ",")
	c.serialize(tree.Left)
	c.serialize(tree.Right)
	return c.Builder.String()
}

func (c *FiveCouple) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	c.Input = strings.Split(data, ",")
	return c.helper()
}

func (c *FiveCouple) helper() *TreeNode {
	if c.Input[0] == "null" {
		c.Input = c.Input[1:]
		return nil
	}
	val, _ := strconv.Atoi(c.Input[0])
	c.Input = c.Input[1:]
	return &TreeNode{
		Val:   val,
		Left:  c.helper(),
		Right: c.helper(),
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
