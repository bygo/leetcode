package tree

func preorderStack(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{}
	for root != nil || 0 < len(stack) {
		for root != nil {
			res = append(res, root.Val)       //前序输出
			stack = append(stack, root.Right) //右节点 入栈
			root = root.Left                  //左节点 遍历
		}
		root = stack[len(stack)-1] //右节点 出栈
		stack = stack[:len(stack)-1]
	}
	return res
}
