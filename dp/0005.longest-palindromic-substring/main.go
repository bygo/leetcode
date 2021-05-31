package main

//Link: https://leetcode-cn.com/problems/longest-palindromic-substring

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	dp := make([][]bool, n)
	var left, right int
	for r := 0; r < n; r++ {
		dp[r] = make([]bool, n)
		dp[r][r] = true // 本身是解

		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) { // 两个解
				dp[l][r] = true

				if right-left < r-l { // 最大长度
					right = r
					left = l
				}
			}
		}
	}

	return s[left:right]
}
