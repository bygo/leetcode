package main

// https://leetcode-cn.com/problems/word-break

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, w := range wordDict {
		m[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for r := 1; r <= n; r++ {
		for l := 0; l < r; l++ {
			if dp[l] && m[s[l:r]] {
				dp[r] = true
				break
			}
		}
	}
	return dp[n]
}
