package main

// https://leetcode-cn.com/problems/count-univalue-subtrees/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 统计同值子树

func countUnivalSubtrees(root *TreeNode) int {
	var cntRes int
	var dfs func(node *TreeNode, parent int) bool
	dfs = func(node *TreeNode, parent int) bool {
		if node == nil {
			return true
		}

		okLeft := dfs(node.Left, node.Val)
		okRight := dfs(node.Right, node.Val)
		if okLeft && okRight {
			cntRes++
			return node.Val == parent
		}
		return false
	}
	dfs(root, 0)
	return cntRes
}
