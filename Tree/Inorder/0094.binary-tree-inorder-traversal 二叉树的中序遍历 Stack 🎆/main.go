package main

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

// ❓ 二叉树的中序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nums []int
	var stack []*TreeNode
	for 0 < len(stack) || root != nil {
		for root != nil {
			// 保存上下文 root包含 right 与 root
			stack = append(stack, root)
			root = root.Left
		}

		// 恢复上下文
		top := len(stack) - 1
		nums = append(nums, stack[top].Val)

		root = stack[top].Right
		// 出栈
		stack = stack[:top]
	}
	return nums
}
