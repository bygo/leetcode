package main

//Link: https://leetcode-cn.com/problems/longest-palindromic-substring

// dp
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)

	var left, right int
	for r := 0; r < n; r++ {
		dp[r] = make([]bool, n)
		//dp[r][r] = true
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) { // 转移
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
		for l := 0; l <= r; l++ {
			dp[l] = s[r] == s[l] && (r-l <= 2 || dp[l+1]) // 压缩需覆盖
			if dp[l] && right-left < r-l {
				left = l
				right = r
			}
		}
	}
	return s[left : right+1]
}
