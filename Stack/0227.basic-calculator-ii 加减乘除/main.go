package main

// https://leetcode-cn.com/problems/basic-calculator-ii

func calculate(s string) int {
	var stack []int
	var i = 0
	var n = len(s)
	var symbol byte = '+'
	for i < n {
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			symbol = s[i]
			i++
		} else if s[i] == ' ' {
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
