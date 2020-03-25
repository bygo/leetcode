package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(node *TreeNode) bool {
	return find(node) != -1
}

func find(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	l := find(node.Left)
	if l == -1 { //剪枝，不平衡时直接返回，
		return -1
	}

	r := find(node.Right)
	if r == -1 { //剪枝，不平衡时直接返回
		return -1
	}

	if math.Abs(l-r) > 1 { //剪枝，不平衡时直接返回
		return -1
	}

	return math.Max(l, r) + 1 //计算深度
}