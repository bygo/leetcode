package main

// https://leetcode-cn.com/problems/perfect-squares

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = n + 1
		for j := 1; j*j <= i; j++ {
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
