package main

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索树中第K小的元素

func kthSmallest(root *TreeNode, k int) int {
	var num int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		k--
		if k == 0 {
			num = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return num
}
