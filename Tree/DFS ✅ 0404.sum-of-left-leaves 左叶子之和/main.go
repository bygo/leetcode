package main

// https://leetcode-cn.com/problems/sum-of-left-leaves/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 左叶子之和

func sumOfLeftLeaves(root *TreeNode) int {
	var sum int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)

		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}
	}
	dfs(root)
	return sum
}
