package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		pre := len(stack) - 1
		root = stack[pre].Right
		stack = stack[:pre]
	}
	return res
}
