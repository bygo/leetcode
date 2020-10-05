package tree

func postorderStack(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		top := len(stack) - 1 //栈顶
		root = stack[top]     //出栈
		stack = stack[:top]
		if root != nil {
			res = append(res, root.Val) //根右左
			stack = append(stack, root.Left, root.Right)
		}
	}

	//反转 变成后序遍历 左右根
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
