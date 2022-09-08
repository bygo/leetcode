package main

// https://leetcode.cn/problems/power-of-two

// ❓ 2的幂

func isPowerOfTwo(n int) bool {
	// 1<<31 的约数
	return 0 < n && 1<<31%n == 0 // TODO
}

func isPowerOfTwo(n int) bool {
	return 0 < n && n&(n-1) == 0 // TODO 最后有效位
}

func isPowerOfTwo(n int) bool {
	// 000010000
	// 111101111 + 1
	return 0 < n && n&-n == n // TODO 最后有效位
}
