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

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var strs []string
	var que = []*TreeNode{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		for _, node := range que[:queL] {
			if node == nil {
				strs = append(strs, "#")
				continue
			}
			strs = append(strs, strconv.Itoa(node.Val))
			que = append(que, node.Left, node.Right)
		}
		que = que[queL:]
	}

	return strings.Join(strs, ",")
}

func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	var strs = strings.Split(data, ",")
	var root = &TreeNode{}
	root.Val, _ = strconv.Atoi(strs[0])
	var que = []*TreeNode{root}
	strs = strs[1:]
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		for _, node := range que[:queL] {
			if strs[0] != "#" {
				numLeft, _ := strconv.Atoi(strs[0])
				node.Left = &TreeNode{Val: numLeft}
				que = append(que, node.Left)
			}

			if strs[1] != "#" {
				numRight, _ := strconv.Atoi(strs[1])
				node.Right = &TreeNode{Val: numRight}
				que = append(que, node.Right)
			}
			strs = strs[2:]
		}
		que = que[queL:]
	}
	return root
}
