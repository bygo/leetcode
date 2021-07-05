package main

// Link: https://leetcode-cn.com/problems/distinct-subsequences

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] // 少一个字符 的路径数
			if s[i-1] == t[j-1] { // 如果相等 加 上次普通路径数
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}
