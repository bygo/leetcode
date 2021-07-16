package main

// https://leetcode-cn.com/problems/longest-valid-parentheses

// 一维
// if **() : f(n) = f(n-2) + 2
// if (*)) : f(n) = f(n-1) + 2 + f(n-f(n-1)-2)
func longestValidParentheses(s string) int {
	var res int
	var n = len(s)
	var f = make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == ')' { // 1. 当前  )
			if s[i-1] == '(' { // 2.前置  (
				if 2 <= i {
					f[i] = f[i-2] + 2
				} else {
					f[i] = 2
				}
			} else if 0 < i-f[i-1] && s[i-f[i-1]-1] == '(' {
				if 2 <= i-f[i-1] {
					f[i] = f[i-1] + f[i-f[i-1]-2] + 2
				} else {
					f[i] = f[i-1] + 2
				}
			}
			if res < f[i] {
				res = f[i]
			}
		}
	}
	return res
}

// map
func longestValidParentheses(s string) int {
	var res int
	var n = len(s)
	var f = map[int]int{}
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				f[i] = f[i-2] + 2
			} else if 0 < i-f[i-1] && s[i-f[i-1]-1] == '(' {
				f[i] = f[i-1] + f[i-f[i-1]-2] + 2
			}
			if res < f[i] {
				res = f[i]
			}
		}
	}
	return res
}
