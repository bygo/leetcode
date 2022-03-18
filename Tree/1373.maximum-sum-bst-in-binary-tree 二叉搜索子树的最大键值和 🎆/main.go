package main

// https://leetcode-cn.com/problems/maximum-sum-bst-in-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索子树的最大键值和

func maxSumBST(root *TreeNode) int {
	var sumMax int
	var dfs func(node *TreeNode) (int, int, bool, int)
	dfs = func(node *TreeNode) (int, int, bool, int) {
		if node == nil {
			return 1<<63 - 1, -1 << 63, true, 0
		}

		minLeft, maxLeft, bLeft, sumLeft := dfs(node.Left)
		minRight, maxRight, bRight, sumRight := dfs(node.Right)

		if node.Val <= maxLeft || minRight <= node.Val || !bLeft || !bRight {
			return 0, 0, false, 0
		}

		sumCur := sumLeft + sumRight + node.Val
		if sumMax < sumCur {
			sumMax = sumCur
		}
		return min(minLeft, node.Val), max(maxRight, node.Val), true, sumCur
	}
	_, _, b, sum := dfs(root)
	if b && sumMax < sum {
	}
	return sumMax
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
