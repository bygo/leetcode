package main

// https://leetcode.cn/problems/binary-tree-preorder-traversal/

// ❓ 前序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var stack = []*TreeNode{}
	var nums []int
	for root != nil {
		for root != nil {
			nums = append(nums, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}

		sTop := len(stack) - 1
		for 0 <= sTop && root == nil {
			root = stack[sTop]
			stack = stack[:sTop]
			sTop--
		}
	}
	return nums
}
