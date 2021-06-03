package main

//Link: https://leetcode-cn.com/problems/wildcard-matching

//s = "adceb"
//p = "*a*b"
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] != '*' {
			break
		}
		dp[0][i] = true
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j] // 不要* 或者 不要字符
			} else if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
