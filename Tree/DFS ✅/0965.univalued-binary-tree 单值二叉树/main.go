package main

// https://leetcode-cn.com/problems/univalued-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 单值二叉树

func isUnivalTree(root *TreeNode) bool {
	var num = root.Val
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		return node == nil || node.Val == num && dfs(node.Left) && dfs(node.Right)
	}
	return dfs(root)
}
