package main

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

// ❓ 前序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nums []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		top := len(stack) - 1
		root = stack[top]
		stack = stack[:top]
		nums = append(nums, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}

		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}

	return nums
}
