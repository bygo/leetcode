package main

type Node struct {
	Val      int
	Children []*Node
}

//Title: Clone N-ary Tree
//Link: https://leetcode-cn.com/problems/clone-n-ary-tree

func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}

	var node = &Node{Val: root.Val}
	for _, c := range root.Children {
		node.Children = append(node.Children, cloneTree(c))
	}
	return node
}

