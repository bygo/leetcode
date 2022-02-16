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
		okLeft := dfs(node.Left)
		if okLeft {
			return true
		}
		okRight := dfs(node.Right)
		if okRight {
			return true
		}
		sumCur -= node.Val
		return false
	}
	return dfs(root)
}
