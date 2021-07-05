package main

// Link: https://leetcode-cn.com/problems/triangle

// 压缩
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	var dp = make([]int, n)

	dp[0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i] = dp[i-1] + triangle[i][i] // 最后一个
		for j := i - 1; 0 < j; j-- {     // 中间
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] = dp[0] + triangle[i][0] // 第一个
	}

	for i := 1; i < n; i++ {
		if dp[i] < dp[0] {
			dp[0] = dp[i]
		}
	}

	return dp[0]
}

// 继续压缩
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	var dp = make([]int, n+1)
	for i := n - 1; 0 <= i; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
