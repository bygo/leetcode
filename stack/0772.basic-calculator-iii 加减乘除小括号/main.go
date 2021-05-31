package main

//Link: https://leetcode-cn.com/problems/basic-calculator-iii

// (2+6*3+5-(3*14/7+2)*5)+3
func calculate(s string) int {
	var stack []int
	var i = 0
	var n = len(s)
	var symbol byte = '+'
	var fuckStack [][]int
	var symbols []byte
	for i < n {
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			symbol = s[i]
			i++
		} else if s[i] == ' ' { //
			i++
		} else if s[i] == '(' { // 保存当前栈
			fuckStack = append(fuckStack, stack)
			symbols = append(symbols, symbol)
			symbol = '+'
			stack = []int{}
			i++
		} else if s[i] == ')' { // 出栈
			var cur int
			for _, v := range stack {
				cur += v
			}

			top := len(fuckStack) - 1

			stack = fuckStack[top]
			fuckStack = fuckStack[:top]

			symbol = symbols[top]
			symbols = symbols[:top]

			switch symbol {
			case '+':
				stack = append(stack, cur)
			case '-':
				stack = append(stack, -cur)
			case '*':
				stack[len(stack)-1] *= cur
			case '/':
				stack[len(stack)-1] /= cur
			}
			i++
		} else {
			var num = 0
			for i < n && '0' <= s[i] && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			switch symbol {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
		}
	}
	var res int
	for _, v := range stack {
		res += v
	}
	return res
}
