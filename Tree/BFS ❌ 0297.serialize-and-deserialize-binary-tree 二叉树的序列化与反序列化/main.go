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
}

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
		for _, q := range que[:queL] {
			if q == nil {
				strs = append(strs, "#")
			} else {
				strs = append(strs, strconv.Itoa(q.Val))
				que = append(que, q.Left, q.Right)
			}
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
		for _, q := range que[:queL] {
			if strs[0] != "#" {
				v, _ := strconv.Atoi(strs[0])
				q.Left = &TreeNode{Val: v}
				que = append(que, q.Left)
			}

			if strs[1] != "#" {
				r, _ := strconv.Atoi(strs[1])
				q.Right = &TreeNode{Val: r}
				que = append(que, q.Right)
			}
			strs = strs[2:]
		}
		que = que[queL:]
	}
	return root
}
