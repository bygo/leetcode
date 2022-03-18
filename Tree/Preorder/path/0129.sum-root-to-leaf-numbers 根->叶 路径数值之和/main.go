package main

// https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 根->叶 路径数值之和

func sumNumbers(root *TreeNode) int {
	var sumTot int
	var sumCur int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		sumCur = sumCur*10 + node.Val
		if node.Left == nil && node.Right == nil {
			sumTot += sumCur
		}
		dfs(node.Left)
		dfs(node.Right)
		sumCur /= 10
	}
	dfs(root)
	return sumTot
}
