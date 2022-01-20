package main

// https://leetcode-cn.com/problems/power-of-two

// ❓ 2的幂

func isPowerOfTwo(n int) bool {
	// BK 只有一个1
	return 0 < n && n&(n-1) == 0
}

func isPowerOfTwo_(n int) bool {
	// 1<<63 的约数
	return 0 < n && 0b10000000000000000000000000000000%n == 0
}
