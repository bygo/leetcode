package main

// https://leetcode-cn.com/problems/path-sum/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 路径总和

func hasPathSum(root *TreeNode, targetSum int) bool {
	var sumCur int
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		sumCur += node.Val
		if node.Left == nil && node.Right == nil && sumCur == targetSum {
			return true
		}
		if dfs(node.Left) || dfs(node.Right) {
			return true
		}

		sumCur -= node.Val
		return false
	}
	return dfs(root)
}
