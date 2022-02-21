package main

// https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索树节点最小距离
// ⚠️ BST

func minDiffInBST(root *TreeNode) int {
	var numDiff = 1<<63 - 1
	var pre *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != nil {
			numCur := node.Val - pre.Val
			if numCur < numDiff {
				numDiff = numCur
			}
		}
		pre = node
		dfs(node.Right)
	}
	dfs(root)
	return numDiff
}

// 0530 https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
