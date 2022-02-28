package main

// https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 从根到叶的二进制数之和

func sumRootToLeaf(root *TreeNode) int {
	var sumRes int
	var dfs func(node *TreeNode, sumCur int)
	dfs = func(node *TreeNode, sumCur int) {
		if node == nil {
			return
		}
		sumCur <<= 1
		sumCur += node.Val
		if node.Left == nil && node.Right == nil {
			sumRes += sumCur
			return
		}
		dfs(node.Left, sumCur)
		dfs(node.Right, sumCur)
	}
	dfs(root, 0)
	return sumRes
}
