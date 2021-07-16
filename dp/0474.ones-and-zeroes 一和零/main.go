package main

// https://leetcode-cn.com/problems/ones-and-zeroes

func findMaxForm(strs []string, m, n int) int {
	length := len(strs)
	dp := make([][][]int, length+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i, s := range strs {
		zeros, ones := get(s)
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				if zeros <= j && ones <= k {
					// 第i个
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zeros][k-ones]+1)
				}
			}
		}
	}
	return dp[length][m][n]
}

// 输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
func findMaxForm(strs []string, m, n int) int {
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeros, ones := get(s)

		for i := m; zeros <= i; i-- {
			for j := n; ones <= j; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func get(s string) (zeros, ones int) {
	for _, v := range s {
		if v == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
