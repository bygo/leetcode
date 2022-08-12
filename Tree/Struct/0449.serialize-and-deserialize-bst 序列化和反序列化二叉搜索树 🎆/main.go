package main

// https://leetcode.cn/problems/serialize-and-deserialize-bst/

import (
	"strconv"
	"strings"
)

// 序列化和反序列化二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	strs []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

func (this *Codec) deserialize(data string) *TreeNode {
	this.strs = strings.Split(data, ",")
	return this.dfs()
}

func (this *Codec) dfs() *TreeNode {
	str := this.strs[0]
	this.strs = this.strs[1:]
	if str == "#" {
		return nil
	}

	num, _ := strconv.Atoi(str)
	root := &TreeNode{
		Val:   num,
		Left:  this.dfs(),
		Right: this.dfs(),
	}
	return root
}
