package main

// Link: https://leetcode-cn.com/problems/basic-calculator

// 从内至外 计算值
func calculate(s string) int {
	var stack []int
	var symbols []int
	var symbol = 1
	var cur int
	var n = len(s)
	var i = 0
	for i < n {
		switch s[i] {
		case '(': // 当前值入栈
			stack = append(stack, cur)
			symbols = append(symbols, symbol)
			symbol = 1
			cur = 0
			i++
		case ')': // 出栈 与 当前 计算
			top := len(stack) - 1
			cur = stack[top] + cur*symbols[top]
			stack = stack[:top]
			symbols = symbols[:top]
			i++
		case '+':
			symbol = 1
			i++
		case '-':
			symbol = -1
			i++
		case ' ':
			i++
		default:
			var num = 0
			for i < n && '0' <= s[i] && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			cur += num * symbol
		}
	}
	return cur
}

// 从左至右 计算符号
func calculate_2(s string) int {
	var res int
	var symbols = []int{1}
	var symbol = 1
	var n = len(s)
	var i = 0
	for i < n {
		switch s[i] {
		case '(':
			symbols = append(symbols, symbol)
			i++
		case ')':
			symbols = symbols[:len(symbols)-1]
			i++
		case '+':
			symbol = symbols[len(symbols)-1]
			i++
		case '-':
			symbol = -symbols[len(symbols)-1]
			i++
		case ' ':
			i++
		default:
			var num = 0
			for i < n && '0' <= s[i] && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			res += num * symbol
		}
	}
	return res
}
