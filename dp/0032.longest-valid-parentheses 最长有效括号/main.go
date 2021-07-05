package main

// Link: https://leetcode-cn.com/problems/longest-valid-parentheses

// f(n) = f(n-2) + 2
// f(n) = f(n-1) + 2 + f(n-f(n-1)-2)
func longestValidParentheses(s string) int {
	var res int
	var dp = make([]int, len(s))
	var n = len(s)
	for i := 1; i < n; i++ {
		if s[i] == ')' { // 1. 当前  )
			if s[i-1] == '(' { // 2.前置  (
				if 2 <= i { // 3.子状态 dp[i-2]
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if 0 < i-dp[i-1] && s[i-dp[i-1]-1] == '(' {
				if 2 <= i-dp[i-1] { // s[i-1] 必是 )
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2 // 4.子状态 dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
