package main

// Link: https://leetcode-cn.com/problems/longest-palindromic-substring

// 双串
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	var left, right int
	for r := 0; r < n; r++ {
		for l := 0; l < n; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if right-left < r-l {
					right = r
					left = l
				}
			}
		}
	}
	return s[left : right+1]
}

// 压缩
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([]bool, n)

	var left, right int
	for r := 0; r < n; r++ {
		for l := 0; l < r; l++ {
			dp[l] = s[r] == s[l] && (r-l <= 2 || dp[l+1])
			if dp[l] && right-left < r-l {
				left = l
				right = r
			}
		}
	}
	return s[left : right+1]
}
