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
		last := len(stack) - 1
		res = append(res, stack[last].Val)
		root = stack[last].Right
		stack = stack[:last]
	}
	return res
}
