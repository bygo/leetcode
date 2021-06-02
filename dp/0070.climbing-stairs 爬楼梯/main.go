package main

//Link: https://leetcode-cn.com/problems/climbing-stairs

func climbStairs(n int) int {
	var a, b, c = 0, 0, 1
	for 0 < n {
		a = b
		b = c
		c = a + b
		n--
	}
	return c
}
// 0 1 1
// 0 1 1 2
// 0 1 1 2 3 5 8 13
