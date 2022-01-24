package main

// https://leetcode-cn.com/problems/power-of-four

func isPowerOfFour(n int) bool {
	// 一个1，且余数为1
	// A. 4-3 = 1
	// B. 剩下1必须继续组成4，才是4的幂
	// A. 4-3 = 1
	// 最后结果必然是 n%3 == 1
	return 0 < n && n&(n-1) == 0 && n%3 == 1
	// ❌ n%4 1 8 12 20
}

func isPowerOfFour_(n int) bool {
	// 一个1，且其他位为0
	return 0 < n && n&(n-1) == 0 && n&0b10101010101010101010101010101010 == 0
}
