package main

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

// ❓ 前序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var nums []int
	var stack []*TreeNode
	for 0 < len(stack) || root != nil {
		for root != nil {
			stack = append(stack, root.Right)
			nums = append(nums, root.Val)
			root = root.Left
		}
		top := len(stack) - 1
		root = stack[top]
		stack = stack[:top]
	}

	return nums
}
