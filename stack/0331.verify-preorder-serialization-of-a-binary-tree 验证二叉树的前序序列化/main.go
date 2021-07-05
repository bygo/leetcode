package main

// Link: https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree

func isValidSerialization(preorder string) bool {
	n := len(preorder)
	stack := []int{1} // 与 root 抵消
	var i int
	for i < n {
		top := len(stack) - 1
		if -1 == top {
			return false
		}
		stack[top]--         // 遍历到一个子节点 父节点计数 -1
		if stack[top] == 0 { // 抵消出栈
			stack = stack[:top]
		}

		if preorder[i] == '#' { //  nil 不需要子节点抵消
			i += 2 // 跳过逗号
		} else {
			stack = append(stack, 2) // 每个节点需要 2个子节点抵消

			for i < n && preorder[i] != ',' { // 移到 逗号
				i++
			}
			i++ // 跳过逗号
		}
	}
	return len(stack) == 0
}

func isValidSerialization(preorder string) bool {
	n := len(preorder)
	counter := 1
	var i int
	for i < n {
		if counter == 0 { // 无人接
			return false
		}
		if preorder[i] == '#' {
			counter--
			i += 2
		} else {
			counter++ // counter = counter - 父1 + 自身2

			for i < n && preorder[i] != ',' {
				i++
			}
			i++
		}
	}
	return counter == 0
}
