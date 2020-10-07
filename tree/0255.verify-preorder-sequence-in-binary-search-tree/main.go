package main

func verifyPreorder(preorder []int) bool {
	stack := make([]int, 0)
	if len(preorder) < 2 {
		return true
	}
	var min = -1 << 63
	stack = append(stack, preorder[0])
	l := len(preorder)
	for i := 1; i < l; i++ {
		for 0 < len(stack) && stack[len(stack)-1] < preorder[i] { //如果出现大值，一直出栈，进入右树
			min = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		if preorder[i] < min { //如果还是不够大，错误的树
			return false
		}

		stack = append(stack, preorder[i]) //如果小值，入栈
	}

	return true
}
