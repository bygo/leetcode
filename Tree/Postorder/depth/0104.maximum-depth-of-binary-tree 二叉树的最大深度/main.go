package main

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的最大深度

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
