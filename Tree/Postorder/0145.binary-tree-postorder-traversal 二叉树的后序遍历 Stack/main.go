package main

// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

// ❓ 二叉树的后序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nums []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		for root != nil {
			nums = append(nums, root.Val)
			stack = append(stack, root.Left)
			root = root.Right
		}
		top := len(stack) - 1
		root = stack[top]
		stack = stack[:top]
	}

	//反转 变成后序遍历 左右根
	lo, hi := 0, len(nums)-1
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
	return nums
}
