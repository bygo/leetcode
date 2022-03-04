package main

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的最小深度

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depMin = 1<<63 - 1
	var depCur int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		depCur++
		if node.Left == nil && node.Right == nil {
			if depCur < depMin {
				depMin = depCur
			}
		}
		dfs(node.Left)
		dfs(node.Right)
		depCur--
	}
	dfs(root)
	return depMin
}
