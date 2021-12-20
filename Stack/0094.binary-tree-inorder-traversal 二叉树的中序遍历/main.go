package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for 0 < len(stack) || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := len(stack) - 1
		res = append(res, stack[top].Val)
		root = stack[top].Right
		stack = stack[:top]
	}
	return res
}
