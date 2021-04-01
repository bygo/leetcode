package _0_11_12

func isValidBST(root *TreeNode) bool {
	var last = &TreeNode{Val: -1 << 63}
	var stack []*TreeNode
	for 0 < len(stack) || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := len(stack) - 1
		if stack[top].Val <= last.Val {
			return false
		}
		last = stack[top]
		root = stack[top].Right
		stack = stack[:top]
	}
	return true
}
