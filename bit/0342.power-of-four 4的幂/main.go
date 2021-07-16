package main

// https://leetcode-cn.com/problems/power-of-four

func isPowerOfFour(n int) bool {
	return 0 < n && n&(n-1) == 0 && n&-1431655766 == 0
}