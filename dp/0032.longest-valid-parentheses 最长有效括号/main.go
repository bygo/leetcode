package main

//Link: https://leetcode-cn.com/problems/longest-valid-parentheses

func longestValidParentheses(s string) int {
	var res int
	var dp = make([]int, len(s))

	for i := 1; i < len(s); i++ {
		if s[i] == ')' { // 只匹配)
			if s[i-1] == '(' { // 刚好是(
				if 2 <= i { // 加上子结果 dp[i-2]
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if 0 < i-dp[i-1] && s[i-dp[i-1]-1] == '(' {
				if 2 <= i-dp[i-1] { // s[i-1] 必是 )
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2 // 加上子结果 dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
