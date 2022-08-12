package main

// https://leetcode.cn/problems/ternary-expression-parser

// 从右往左 遇到 ? 结算
func parseTernary(expression string) string {
	var stack []byte
	i := len(expression) - 1
	for -1 < i {
		switch expression[i] {
		case '?':
			i--
			top := len(stack) - 1
			t := stack[top]
			f := stack[top-1]
			stack = stack[:top-1]
			if expression[i] == 'T' {
				stack = append(stack, t)
			} else {
				stack = append(stack, f)
			}
		case ':':
		default:
			stack = append(stack, expression[i])
		}
		i--
	}
	return string(stack[0])
}
