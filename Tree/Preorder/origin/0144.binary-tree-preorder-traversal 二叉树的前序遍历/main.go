package main

// https://leetcode.cn/problems/binary-tree-preorder-traversal/

// ❓ 前序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var nums []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nums = append(nums, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return nums
}
