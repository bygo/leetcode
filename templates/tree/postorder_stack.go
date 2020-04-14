package tree

func postorderStack(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		if root != nil {
			//根右左
			res = append(res, root.Val)
			stack = append(stack, root.Left)  //左节点，因为先进 所以后出
			stack = append(stack, root.Right) //右节点，因为后进 所以先出
		}
		index := len(stack) - 1 //栈顶
		root = stack[index]     //出栈
		stack = stack[:index]
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
