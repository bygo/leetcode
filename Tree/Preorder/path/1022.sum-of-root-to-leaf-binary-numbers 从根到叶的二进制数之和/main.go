package main

// https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 从根到叶的二进制数之和

func sumRootToLeaf(root *TreeNode) int {
	var sumRes int
	var sumCur int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		sumCur <<= 1
		sumCur += node.Val
		defer func() {
			sumCur -= node.Val
			sumCur >>= 1
		}()
		if node.Left == nil && node.Right == nil {
			sumRes += sumCur
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sumRes
}
