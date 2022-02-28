package main

// https://leetcode-cn.com/problems/maximum-difference-between-node-and-ancestor/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 节点与节点 与其祖先之间的最大差值

func maxAncestorDiff(root *TreeNode) int {
	var numDiffMax int
	var dfs func(node *TreeNode, min, max int)
	dfs = func(node *TreeNode, min, max int) {
		if node == nil {
			return
		}
		if node.Val < min {
			min = node.Val
		} else if max < node.Val {
			max = node.Val
		}

		numDiffCur := max - min
		if numDiffMax < numDiffCur {
			numDiffMax = numDiffCur
		}
		dfs(node.Left, min, max)
		dfs(node.Right, min, max)
	}
	dfs(root, root.Val, root.Val)
	return numDiffMax
}
