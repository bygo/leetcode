package main

// https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree



func verifyPreorder(preorder []int) bool {
	if len(preorder) < 2 {
		return true
	}

	var root = -1 << 63 //第一次判断
	var stack = []int{}

	for _, v := range preorder {
		top := len(stack) - 1
		// 进入左树 单调递减栈
		for 0 <= top && stack[top] < v { //有比栈顶 大的值，出栈 准备进入子右树
			root = stack[top]
			stack = stack[:top]
			top--
		}
		// 进入右树

		// 我们在右树，
		// 正确的值应该  最后一个root <= v
		if v < root {
			return false
		}

		stack = append(stack, v) //递减值 ，入栈，进入左树
	}

	return true
}
