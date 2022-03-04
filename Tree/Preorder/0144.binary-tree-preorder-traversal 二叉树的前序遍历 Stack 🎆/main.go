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
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		for root != nil {
			nums = append(nums, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}
		top := len(stack) - 1
		root = stack[top]
		stack = stack[:top]
	}

	return nums
}
