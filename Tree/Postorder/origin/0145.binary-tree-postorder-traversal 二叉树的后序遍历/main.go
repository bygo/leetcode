package main

// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

// ❓ 二叉树的后序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var nums []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		nums = append(nums, node.Val)
	}
	dfs(root)
	return nums
}
