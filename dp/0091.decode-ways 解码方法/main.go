package main

// Link: https://leetcode-cn.com/problems/decode-ways

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}
	for i := 2; i <= n; i++ {
		if s[i-1] != '0' { // 当前不是0 组合 1 种
			dp[i] += dp[i-1]
		}

		if s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 { // 前面不是0 就可以组合 2 种
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

// 压缩
func numDecodings(s string) int {
	n := len(s)
	var a, b, c = 1, 0, 0
	if s[0] != '0' {
		b = 1
		c = 1
	}
	for i := 1; i < n; i++ {
		c = 0
		if s[i] != '0' {
			c += b
		}
		if s[i-1] != '0' && ((s[i-1]-'0')*10+(s[i]-'0') <= 26) {
			c += a
		}
		a, b = b, c
	}
	return c
}
