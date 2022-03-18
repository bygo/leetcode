package main

// https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 两数之和 = target

func findTarget(root *TreeNode, target int) bool {
	numMp := map[int]struct{}{}
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		_, ok := numMp[target-root.Val]
		numMp[root.Val] = struct{}{}
		return ok || dfs(root.Left) || dfs(root.Right)
	}
	return dfs(root)
}
