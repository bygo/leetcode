package main

// https://leetcode.cn/problems/reverse-substrings-between-each-pair-of-parentheses

func reverseParentheses(s string) string {
	var stack [][]byte
	var res []byte
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, res)
			res = []byte{}
		} else if s[i] == ')' {
			top := len(stack) - 1
			l := len(res)
			for j := 0; j < l/2; j++ {
				res[j], res[l-j-1] = res[l-j-1], res[j]
			}
			res = append(stack[top], res...)
			stack = stack[:top]
		} else if 'a' <= s[i] && s[i] <= 'z' {
			res = append(res, s[i])
		}
	}

	return string(res)
}
