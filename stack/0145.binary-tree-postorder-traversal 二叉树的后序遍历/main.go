package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Link: https://leetcode-cn.com/problems/binary-tree-postorder-traversal

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		top := len(stack) - 1
		root = stack[top]
		stack = stack[:top]
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Left, root.Right)
		}
	}
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
