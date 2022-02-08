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
	var stack []*TreeNode

	for root != nil || 0 < len(stack) {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		top := len(stack) - 1
		nums = append(nums, stack[top].Val)
		root = stack[top].Right
		stack = stack[:top]
	}
	return nums
}
