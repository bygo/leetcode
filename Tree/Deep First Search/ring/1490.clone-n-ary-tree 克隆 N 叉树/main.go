package main

// https://leetcode.cn/problems/clone-n-ary-tree

type Node struct {
	Val      int
	Children []*Node
}

// ❓ 克隆 N 叉树

func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}

	var node = &Node{Val: root.Val}
	for _, child := range root.Children {
		node.Children = append(node.Children, cloneTree(child))
	}
	return node
}
