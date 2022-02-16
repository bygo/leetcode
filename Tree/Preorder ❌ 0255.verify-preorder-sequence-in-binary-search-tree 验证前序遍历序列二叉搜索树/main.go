package main

// https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree/

// ❓ 验证前序遍历序列二叉搜索树

func verifyPreorder(preorder []int) bool {
	if len(preorder) < 2 {
		return true
	}

	var root = -1 << 63 //第一次判断
	var stack = []int{}

	for _, num := range preorder {
		top := len(stack) - 1
		for 0 <= top && stack[top] < num { //大值，出栈，准备进入右树
			root = stack[top]
			stack = stack[:top]
			top--
		}

		if num < root {
			//左树最后的根节点大于当前节点=错误的树
			return false
		}

		stack = append(stack, num) //小值，入栈，准备进入左树
	}

	return true
}
