package main

// https://leetcode.cn/problems/climbing-stairs

// 一维
// f(n) = f(n-1) + f(n-2)

func climbStairs(n int) int {
	f := make([]int, n+1)
	f[0] = 1
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-2] + f[i-1]
	}
	return f[n]
}

func climbStairs_(n int) int {
	x, y := 1, 1
	for 1 <= n {
		y, x = x+y, y
		n--
	}
	return x
}
