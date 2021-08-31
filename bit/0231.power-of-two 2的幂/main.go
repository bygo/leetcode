package main

// https://leetcode-cn.com/problems/power-of-two

func isPowerOfTwo(n int) bool {
	return 0 < n && n&(n-1) == 0
}
