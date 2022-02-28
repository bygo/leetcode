package main

// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的序列化与反序列化

type Codec struct {
	strs []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + c.serialize(root.Left) + "," + c.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	c.strs = strings.Split(data, ",")
	return c.dfs()
}

func (c *Codec) dfs() *TreeNode {
	str := c.strs[0]
	c.strs = c.strs[1:]
	if str == "#" {
		return nil
	}

	num, _ := strconv.Atoi(str)
	root := &TreeNode{
		Val:   num,
		Left:  c.dfs(),
		Right: c.dfs(),
	}
	return root
}
