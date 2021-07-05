package main

// Link: https://leetcode-cn.com/problems/fibonacci-number

// 0 1 1
func fib(n int) int {
	var dp = make([]int, n+2)
	dp[0] = 0
	dp[1] = 1
	var i = 2
	for i <= n {
		dp[i] = dp[i-1] + dp[i-2]
		i++
	}
	return dp[n]
}

// 压缩
func fib(n int) int {
	var a, b = 0, 1
	for 0 < n {
		b, a = (a+b)%1000000007, b
		n--
	}
	return a
}
