package main

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树中的最大路径和

func maxPathSum(root *TreeNode) int {
	var sumMax = -1 << 63
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := max(dfs(node.Left), 0)
		sumRight := max(dfs(node.Right), 0)
		sumCur := sumLeft + sumRight + node.Val
		sumMax = max(sumMax, sumCur)
		return max(sumLeft, sumRight) + node.Val
	}
	dfs(root)
	return sumMax
}

func max(nums ...int) int {
	for _, num := range nums[1:] {
		if nums[0] < num {
			nums[0] = num
		}
	}
	return nums[0]
}
