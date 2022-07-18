package main

// https://leetcode.cn/problems/binary-tree-tilt/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的坡度
// ⚠️ 左右树绝对值差 的和

func findTilt(root *TreeNode) int {
	var tiltTot int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)
		tiltTot += abs(sumLeft - sumRight)
		return node.Val + sumLeft + sumRight
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
