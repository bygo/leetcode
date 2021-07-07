package main

// Link: https://leetcode-cn.com/problems/fibonacci-number

// 0 1 1
func climbStairs(n int) int {
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// 压缩
func fib(n int) int {
	var x, y = 0, 1
	for i := 1; i <= n; i++ {
		y, x = (x+y)%1000000007, y
	}
	return x
}
