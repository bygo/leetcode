package tree

func preorderStack(root *TreeNode) {
	var stack = []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			println(root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		pre := len(stack) - 1
		root = stack[pre].Right
		stack = stack[:pre]
	}
}
