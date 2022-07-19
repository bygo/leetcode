package main

// https://leetcode.cn/problems/power-of-two

// ❓ 2的幂

func isPowerOfTwo(n int) bool {
	// 1<<31 的约数
	return 0 < n && 1<<31%n == 0
}

func isPowerOfTwo(n int) bool {
	// BK 只有一个1
	return 0 < n && n&(n-1) == 0
}

func isPowerOfTwo(n int) bool {
	return 0 < n && n&-n == n
}
