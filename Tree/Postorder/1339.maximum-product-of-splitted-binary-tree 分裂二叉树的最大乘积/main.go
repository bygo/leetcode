package main

// https://leetcode-cn.com/problems/maximum-product-of-splitted-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 分裂二叉树的最大乘积

func maxProduct(root *TreeNode) int {
	var sums []int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)

		sumCur := sumLeft + sumRight + node.Val
		sums = append(sums, sumCur)
		return sumCur
	}

	sumTot := dfs(root)
	proMax := 0
	for _, sum := range sums {
		proCur := sum * (sumTot - sum)
		if proMax < proCur {
			proMax = proCur
		}
	}
	return proMax % 1000000007
}
