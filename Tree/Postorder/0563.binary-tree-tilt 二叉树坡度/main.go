package main

// https://leetcode-cn.com/problems/binary-tree-tilt/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的坡度

func findTilt(root *TreeNode) int {
	var tiltTot int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		tiltLeft := dfs(node.Left)
		tiltRight := dfs(node.Right)
		tiltTot += abs(tiltLeft - tiltRight)
		return node.Val + tiltLeft + tiltRight
	}
	dfs(root)
	return tiltTot
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
