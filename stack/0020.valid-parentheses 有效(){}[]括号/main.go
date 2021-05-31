package main

//Link:https://leetcode-cn.com/problems/valid-parentheses/

// 1.map 镜像
// 2.相同出栈
func isValid(s string) bool {
	m := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	var stack []rune
	for _, v := range s {
		if mirror, ok := m[v]; ok { // 加入 镜像
			stack = append(stack, mirror)
		} else if 0 < len(stack) && stack[len(stack)-1] == v { // 存在时抵消
			stack = stack[:len(stack)-1]
		} else {
			return false // 抵消失败
		}
	}
	return len(stack) == 0 // 完全抵消
}
