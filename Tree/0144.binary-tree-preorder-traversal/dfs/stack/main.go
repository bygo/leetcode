package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
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
