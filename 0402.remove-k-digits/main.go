package main

//Link: https://leetcode-cn.com/problems/remove-k-digits

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		v := num[i]
		for 0 < k && 0 < len(stack) && v < stack[len(stack)-1] { // 单调递增栈
			stack = stack[:len(stack)-1] //栈内大值 pop 等于删除
			k--
		}
		stack = append(stack, v) // push
	}
	stack = stack[:len(stack)-k] // 完全递增 直接删除尾部
	for 1 < len(stack) {
		if stack[0] != '0' {
			break
		}
		stack = stack[1:]
	}
	if 0 == len(stack) {
		return "0"
	}
	return string(stack)
}
