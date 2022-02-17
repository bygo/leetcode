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

type Codec struct{}

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
		for idx, node := range que[:queL] {
			if idx != 0 {
				strs = append(strs, "#")
			}
			for idxChild, child := range node.Children {
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
	var zero, parent = &Node{}, &Node{}
	var que = []*Node{zero}

	depsNodes := strings.Split(data, "d")
	for _, nodes := range depsNodes {
		queL := len(que)
		node := strings.Split(nodes, "#")
		for idx, child := range node {
			parent = que[idx]
			nodesChild := strings.Split(child, ",")
			for _, nodeChild := range nodesChild {
				if nodeChild == "" {
					continue
				}
				num, _ := strconv.Atoi(nodeChild)
				nodeChild := &Node{Val: num}
				parent.Children = append(parent.Children, nodeChild)
				que = append(que, nodeChild)
			}
		}
		que = que[queL:]
	}

	return zero.Children[0]
}
