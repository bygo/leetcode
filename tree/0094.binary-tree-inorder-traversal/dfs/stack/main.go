package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		i := len(stack) - 1
		root, stack = stack[i], stack[:i]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
