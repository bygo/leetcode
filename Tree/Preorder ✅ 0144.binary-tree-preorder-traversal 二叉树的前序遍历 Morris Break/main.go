package main

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

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
		nums = append(nums, root.Val)
		if root.Left == nil {
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil {
				max = max.Right
			}

			max.Right = root.Right
			root, root.Left = root.Left, nil
			//root.Right, max.Right = root.Left, root.Right
			//root.Left = nil
		}
	}
	return nums
}
