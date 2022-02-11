package main

// https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 两数之和 IV - 输入 BST

func findTarget(root *TreeNode, k int) bool {
	numMp := map[int]struct{}{}
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		_, ok := numMp[k-root.Val]
		numMp[root.Val] = struct{}{}
		return ok || dfs(root.Left) || dfs(root.Right)
	}
	return dfs(root)
}
