package main

// https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree/

// ❓ 验证前序遍历序列二叉搜索树

func verifyPreorder(preorder []int) bool {
	if len(preorder) < 2 {
		return true
	}

	var stack = []int{preorder[0]}
	// zero 占位节点
	var root int
	for _, num := range preorder[1:] {
		top := len(stack) - 1
		for 0 <= top && stack[top] < num { // 比栈顶大的值，出栈，准备进入右树, 改变根节点
			root = stack[top]
			stack = stack[:top]
			top--
		}

		// 出现 比 最后根节点小的值
		if num < root {
			return false
		}

		// 单调递减栈
		// 比栈顶小的值，入栈，进入左树
		// min ~ root
		stack = append(stack, num)
	}

	return true
}
