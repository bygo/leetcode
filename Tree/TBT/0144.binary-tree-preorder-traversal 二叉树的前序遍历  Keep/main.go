package main

// https://leetcode.cn/problems/binary-tree-preorder-traversal/

// ❓ 前序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var max *TreeNode
	var nums []int
	for root != nil {
		if root.Left == nil {
			nums = append(nums, root.Val)
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil && max.Right != root {
				max = max.Right
			}

			if max.Right == nil {
				nums = append(nums, root.Val)
				max.Right = root.Right
				root = root.Left
			} else {
				// 已输出
				root = root.Right
				max.Right = nil
			}
		}
	}
	return nums
}
