package main

// https://leetcode-cn.com/problems/recover-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 恢复二叉搜索树

func recoverTree(root *TreeNode) {
	var pre, first, second *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != nil && node.Val <= pre.Val {
			second = node
			if first != nil {
				return
			}
			first = pre
		}
		pre = node
		dfs(node.Right)
	}
	dfs(root)
	first.Val, second.Val = second.Val, first.Val
}
