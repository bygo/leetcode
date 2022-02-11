package main

// https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树中第二小的节点

func findSecondMinimumValue(root *TreeNode) int {
	numMin := root.Val
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		if numMin < node.Val { // 根节点一定最小，如果出现大数，即是答案
			return node.Val
		}

		numL := dfs(node.Left)
		numR := dfs(node.Right)
		if numL == -1 {
			return numR
		}
		if numR == -1 {
			return numL
		}
		return min(numL, numR)
	}
	return dfs(root)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
