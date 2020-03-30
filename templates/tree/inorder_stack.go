package tree

func inorderStack(root *TreeNode) []int {
	var res []int
	stack := []*TreeNode{}
	for root != nil || 0 < len(stack) {
		for root != nil {
			stack = append(stack, root) //根节点入栈
			root = root.Left            //左节点遍历
		}

		res = append(res, stack[len(stack)-1].Val) //中序输出
		root = stack[len(stack)-1].Right           //右节点
		stack = stack[:len(stack)-1]
	}

	return res
}
