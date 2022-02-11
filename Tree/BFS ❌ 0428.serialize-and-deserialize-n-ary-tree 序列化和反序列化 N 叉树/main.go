package main

// https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree/

import (
	"strconv"
	"strings"
)

type Node struct {
	Val      int
	Children []*Node
}

// ❓ 序列化和反序列化 N 叉树

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (c *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}

	var que = []*Node{root}
	var strs = []string{strconv.Itoa(root.Val)}
	for {
		var queL = len(que)
		if queL == 0 {
			break
		}
		strs = append(strs, "d")
		for idx, q := range que[:queL] {
			if idx != 0 {
				strs = append(strs, "#")
			}
			for idxChild, child := range q.Children {
				if idxChild != 0 {
					strs = append(strs, ",")
				}
				que = append(que, child)
				strs = append(strs, strconv.Itoa(child.Val))
			}
		}
		que = que[queL:]
	}

	return strings.Join(strs, "")
}

func (c *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	var pre, parent = &Node{}, &Node{}
	var que = []*Node{pre}
	for _, strsDep := range strings.Split(data, "d") {
		queL := len(que)
		for dep, str := range strings.Split(strsDep, "#") {
			parent = que[dep]
			for _, q := range strings.Split(str, ",") {
				if q != "" {
					v, _ := strconv.Atoi(q)
					child := &Node{Val: v}
					parent.Children = append(parent.Children, child)
					que = append(que, child)
				}
			}
		}
		que = que[queL:]
	}

	return pre.Children[0]
}
