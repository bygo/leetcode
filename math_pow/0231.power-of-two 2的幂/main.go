package main

// https://leetcode-cn.com/problems/power-of-two

func isPowerOfTwo(n int) bool {
	return 0 < n && n&(n-1) == 0
}

func isPowerOfTwo(n int) bool {
	return 0 < n && 0b10000000000000000000000000000000%n == 0
}
