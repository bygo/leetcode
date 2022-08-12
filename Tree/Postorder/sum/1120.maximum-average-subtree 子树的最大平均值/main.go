package main

// https://leetcode.cn/problems/maximum-average-subtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 子树的最大平均值

func maximumAverageSubtree(root *TreeNode) float64 {
	var numAvg float64 = -1 << 63
	var dfs func(node *TreeNode) (float64, float64)
	dfs = func(node *TreeNode) (float64, float64) {
		if node == nil {
			return 0, 0
		}

		numLeft, cntLeft := dfs(node.Left)
		numRight, cntRight := dfs(node.Right)

		numCur := numLeft + numRight + float64(node.Val)
		cntCur := cntLeft + cntRight + 1
		numAvgCur := numCur / cntCur
		if numAvg < numAvgCur {
			numAvg = numAvgCur
		}
		return numCur, cntCur
	}
	dfs(root)
	return numAvg
}
