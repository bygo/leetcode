package main

// https://leetcode.cn/problems/diameter-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的直径

func diameterOfBinaryTree(root *TreeNode) int {
	var cntMax int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		cntL := dfs(node.Left)
		cntR := dfs(node.Right)
		cntCur := max(cntL, cntR)
		cntMax = max(cntMax, cntL+cntR)
		return cntCur + 1
	}
	dfs(root)
	return cntMax
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
