package tree

func preorderStack(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		top := len(stack) - 1 //栈顶
		root = stack[top]     //出栈
		stack = stack[:top]
		if root != nil {
			res = append(res, root.Val) //根左右
			stack = append(stack, root.Right, root.Left)
		}
	}
	return res
}
