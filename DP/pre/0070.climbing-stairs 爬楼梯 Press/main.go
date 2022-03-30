package main

// https://leetcode-cn.com/problems/climbing-stairs

// 一维
// f(n) = f(n-1) + f(n-2)

// 压缩
func climbStairs(n int) int {
	x, y := 1, 1
	for i := 2; i <= n; i++ {
		y, x = x+y, y
	}
	return y
}
