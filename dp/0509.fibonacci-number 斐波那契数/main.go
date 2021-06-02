package main

//Link: https://leetcode-cn.com/problems/fibonacci-number

func fib(n int) int {
	var a, b, c = 0, 0, 1
	for 0 < n {
		a = b
		b = c
		c = (a + b) % 1000000007
		n--
	}
	return b
}

// 0->1
// 0->2

// 0 1 2 3 5 8 13