package main

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

// ❓ 二叉树的中序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var nums []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			nums = append(nums, node.Val)
			dfs(node.Right)
		}
	}
	dfs(root)
	return nums
}
