package tree

func postorderStack(root *TreeNode) []int {
	var res = []int{}
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		if root.Left != nil {
			stack = append(stack, root.Left)
		}

		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}

	l := len(res) - 1
	for i := 0; i < l/2; i++ {
		res[i], res[l-i] = res[l-i], res[i]
	}
	return res
}
