package main

func verifyPreorder(preorder []int) bool {
	if len(preorder) < 2 {
		return true
	}

	var root = -1 << 63 //第一次判断
	var stack = []int{}

	for _, v := range preorder {
		top := len(stack) - 1
		for 0 <= top && stack[top] < v { //大值，出栈，准备进入右树
			root = stack[top]
			stack = stack[:top]
			top--
		}

		if v < root {
			//左树最后的根节点大于当前节点=错误的树
			return false
		}

		stack = append(stack, v) //小值，入栈，准备进入左树
	}

	return true
}
