package main

// https://leetcode.cn/problems/power-of-four

// ❓ 4的幂

func isPowerOfFour(n int) bool {
	return 0 < n && n&(n-1) == 0 && n&0b10101010101010101010101010101010 == 0
}

func isPowerOfFour(n int) bool {
	return 0 < n && n&(n-1) == 0 && n%3 == 1 // TODO
	// ✅ 1 4 16 64 256
	// ❌ n%4==  8 32 128 512
}
