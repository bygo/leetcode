package main

// Link: https://leetcode-cn.com/problems/interleaving-string

// i j 能组成p
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if (l1 + l2) != l3 {
		return false
	}
	dp := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]bool, l2+1)
	}
	dp[0][0] = true
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			p := i + j - 1
			dp[i][j] = 0 < i && dp[i-1][j] && s1[i-1] == s3[p] ||
				0 < j && dp[i][j-1] && s2[j-1] == s3[p] ||
				dp[i][j]
		}
	}
	return dp[l1][l2]
}

// 压缩 dp[j] 表示 当前i下
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if (l1 + l2) != l3 {
		return false
	}
	dp := make([]bool, l2+1)
	dp[0] = true
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			p := i + j - 1
			if 0 < i {
				dp[j] = dp[j] && s1[i-1] == s3[p] // i增加后 dp[j] 是否依然有效
			}

			if 0 < j { // 如果dp[j] 有效 或 可以通过增加j 使之有效
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[p] //
			}
		}
	}
	return dp[l2]
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	dp := make([]bool, m+1)
	dp[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return dp[m]
}
