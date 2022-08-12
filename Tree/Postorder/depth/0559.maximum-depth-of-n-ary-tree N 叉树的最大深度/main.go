package main

// https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/

type Node struct {
	Val      int
	Children []*Node
}

// ❓ N 叉树的最大深度

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var dep int
	for _, child := range root.Children {
		dep = max(dep, maxDepth(child))
	}

	return dep + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
