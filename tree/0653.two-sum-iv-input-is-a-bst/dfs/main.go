package main

// https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := map[int]struct{}{}
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root != nil {
			_, ok := m[k-root.Val]
			m[root.Val] = struct{}{}
			return ok || dfs(root.Left) || dfs(root.Right)
		}
		return false
	}
	return dfs(root)
}
