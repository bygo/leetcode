package main

// https://leetcode-cn.com/problems/longest-zigzag-path-in-a-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树中的最长交错路径

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var cntMax int
	var dfs func(node *TreeNode, cntCur int, parent int)
	dfs = func(node *TreeNode, cntCur int, parent int) {
		if node == nil {
			return
		}
		if cntMax < cntCur {
			cntMax = cntCur
		}
		if parent == -1 {
			dfs(node.Left, 1, -1)
			dfs(node.Right, cntCur+1, 1)
		} else {
			dfs(node.Left, cntCur+1, -1)
			dfs(node.Right, 1, 1)
		}
	}
	dfs(root.Left, 1, -1)
	dfs(root.Right, 1, 1)
	return cntMax
}
