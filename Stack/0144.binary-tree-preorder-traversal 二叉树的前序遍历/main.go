package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal

func preorderTraversal(root *TreeNode) []int {
	var stack = []*TreeNode{root}
	var res = []int{}
	for 0 < len(stack) {
		top := len(stack) - 1
		root = stack[top]
		stack = stack[:top]
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Right, root.Left)
		}
	}
	return res
}
