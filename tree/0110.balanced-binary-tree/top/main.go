package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(node *TreeNode) bool {
	return node == nil || isBalanced(node.Left) &&
		math.Abs(height(node.Left)-height(node.Right)) < 2 && //两边最大深度差  大于2
		isBalanced(node.Right)
}

//计算节点最大深度
func height(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(height(node.Left), height(node.Right)) + 1
}

// 计算左右树深度->对比深度
