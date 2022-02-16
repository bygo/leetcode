package main

// https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 求根节点到叶节点数字之和

func sumNumbers(root *TreeNode) int {
	var sumTot int
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum = sum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			sumTot += sum
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, 0)
	return sumTot
}
