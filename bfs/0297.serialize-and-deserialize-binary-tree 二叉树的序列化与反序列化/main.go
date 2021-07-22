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

// serialize and deserialize binary tree
// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var res []string
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		for _, q := range queue[:cnt] {
			if q != nil {
				res = append(res, strconv.Itoa(q.Val))
				queue = append(queue, q.Left, q.Right)
			} else {
				res = append(res, "#")
			}
		}
		queue = queue[cnt:]
	}

	return strings.Join(res, ",")
}

func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	var res = strings.Split(data, ",")
	var root = &TreeNode{}
	root.Val, _ = strconv.Atoi(res[0])
	var queue = []*TreeNode{root}
	res = res[1:]
	for 0 < len(queue) {
		var cnt = len(queue)
		for _, q := range queue[:cnt] {
			if res[0] != "#" {
				l, _ := strconv.Atoi(res[0])
				q.Left = &TreeNode{Val: l}
				queue = append(queue, q.Left)
			}

			if res[1] != "#" {
				r, _ := strconv.Atoi(res[1])
				q.Right = &TreeNode{Val: r}
				queue = append(queue, q.Right)
			}
			res = res[2:]
		}
		queue = queue[cnt:]
	}
	return root
}
