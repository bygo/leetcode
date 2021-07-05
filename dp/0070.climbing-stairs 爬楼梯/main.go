package main

// Link: https://leetcode-cn.com/problems/climbing-stairs

// dp
func climbStairs(n int) int {
	var dp = make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	var i = 2
	for i <= n {
		dp[i] = dp[i-1] + dp[i-2]
		i++
	}
	return dp[n]
}

// 压缩
func climbStairs(n int) int {
	var x, y = 0, 1
	for 0 < n {
		y, x = x+y, y
		n--
	}
	return y
}

// 1 1 2
