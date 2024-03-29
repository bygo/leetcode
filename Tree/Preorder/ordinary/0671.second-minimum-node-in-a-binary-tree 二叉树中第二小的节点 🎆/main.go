package main

// https://leetcode.cn/problems/second-minimum-node-in-a-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树中第二小的节点
// ⚠️ 子节点为 0或2 个
// ⚠️ 父节点 等于 最小子节点

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
		if numL < numR {
			return numL
		}
		return numR
	}
	return dfs(root)
}
