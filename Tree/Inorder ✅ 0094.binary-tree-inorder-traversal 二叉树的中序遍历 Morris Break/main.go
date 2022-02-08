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
	var max *TreeNode

	for root != nil {
		if root.Left == nil {
			nums = append(nums, root.Val)
			root = root.Right
		} else {
			//寻找左树最大节点
			max = root.Left
			for max.Right != nil {
				max = max.Right
			}

			//最大节点指向root
			max.Right = root

			//root = root.Left :移动root到root.left
			//root.Left = nil  :砍左子树，避免下一次遍历到root时，再进入到左子树
			root, root.Left = root.Left, nil
		}
	}
	return nums
}
