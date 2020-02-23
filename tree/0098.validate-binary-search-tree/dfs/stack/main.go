package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var last *TreeNode
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		pre := len(stack) - 1
		if last != nil && last.Val >= stack[pre].Val {
			return false
		}
		last = stack[pre]
		root = stack[pre].Right
		stack = stack[:pre]
	}

	return true
}
