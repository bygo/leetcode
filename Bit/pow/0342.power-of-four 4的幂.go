package main

// https://leetcode-cn.com/problems/power-of-four

// ❓ 4的幂

func isPowerOfFour(n int) bool {
	// 一个1，且其他位为0
	return 0 < n && n&(n-1) == 0 && n&0b10101010101010101010101010101010 == 0
}

func isPowerOfFour(n int) bool {
	// 一个1，且余数为1
	return 0 < n && n&(n-1) == 0 && n%3 == 1
	// ❌ n%4 1 8 12 20
}
