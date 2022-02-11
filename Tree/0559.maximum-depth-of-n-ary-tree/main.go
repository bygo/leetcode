package main

type Node struct {
	Val      int
	Children []*Node
}

var max int

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var level int
	for _, n := range root.Children {
		level = maxInt(level, maxDepth(n))
	}

	return level + 1
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
