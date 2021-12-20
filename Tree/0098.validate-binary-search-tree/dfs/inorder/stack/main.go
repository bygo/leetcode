package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var last *TreeNode
	var stack []*TreeNode
	for 0 < len(stack) || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		top := len(stack) - 1
		if last != nil && stack[top].Val < last.Val {
			return false
		}
		last = stack[top]
		root = stack[top].Right
		stack = stack[:top]
	}
	return true
}
