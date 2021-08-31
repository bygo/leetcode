package main

// https://leetcode-cn.com/problems/power-of-four

func isPowerOfFour(n int) bool {
	return 0 < n && n&(n-1) == 0 && n&0b10101010101010101010101010101010 == 0
}

//func isPowerOfFour(n int) bool {
//	return 0 < n && n&(n-1) == 0 && n%3 == 1
//}
