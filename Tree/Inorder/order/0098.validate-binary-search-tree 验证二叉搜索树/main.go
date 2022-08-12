package main

// https://leetcode.cn/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 验证二叉搜索树

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
			return false
		}
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		return dfs(node.Right)
	}
	return dfs(root)
}
