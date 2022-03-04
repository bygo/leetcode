package main

// https://leetcode-cn.com/problems/distribute-coins-in-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 在二叉树中分配硬币

func distributeCoins(root *TreeNode) int {
	var cnt int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		cntL := dfs(node.Left)
		cntR := dfs(node.Right)
		cnt += abs(cntL) + abs(cntR)
		return cntL + cntR + node.Val - 1
	}
	dfs(root)
	return cnt
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}
