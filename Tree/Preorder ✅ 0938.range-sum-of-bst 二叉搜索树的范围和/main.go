package main

// https://leetcode-cn.com/problems/range-sum-of-bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索树的范围和
// ⚠️ BST

func rangeSumBST(root *TreeNode, l int, r int) int {
	var sum int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if l <= node.Val && node.Val <= r {
			sum += node.Val
			dfs(node.Left)
			dfs(node.Right)
		} else if node.Val < l {
			dfs(node.Right)
		} else if r < node.Val {
			dfs(node.Left)
		}
	}
	dfs(root)
	return sum
}
