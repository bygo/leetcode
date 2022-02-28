package main

// https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树最长连续 `父-子` 序列

func longestConsecutive(root *TreeNode) int {
	var cntMax = 1
	var dfs func(node *TreeNode, parent int, cntCur int)
	dfs = func(node *TreeNode, parent int, cntCur int) {
		if node == nil {
			return
		}
		if parent+1 != node.Val {
			cntCur = 1
		}
		if cntMax < cntCur {
			cntMax = cntCur
		}
		dfs(node.Left, node.Val, cntCur+1)
		dfs(node.Right, node.Val, cntCur+1)
	}
	dfs(root, root.Val, 1)
	return cntMax
}
