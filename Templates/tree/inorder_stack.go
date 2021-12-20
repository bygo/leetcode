package tree

//左根右
func inorderStack(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for 0 < len(stack) || root != nil { //root != nil 只为了第一次root判断，必须放最后
		for root != nil {
			stack = append(stack, root) //入栈
			root = root.Left            //移至最左
		}

		top := len(stack) - 1             //栈顶
		res = append(res, stack[top].Val) //中序遍历
		root = stack[top].Right           //右节点会进入for循环，如果=nil，继续出栈
		stack = stack[:top]               //出栈
	}
	return res
}
