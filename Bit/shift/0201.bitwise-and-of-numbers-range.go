package main

// https://leetcode.cn/problems/bitwise-and-of-numbers-range

func rangeBitwiseAnd(lo int, hi int) int {
	shift := 0
	for lo < hi {
		// TODO 共同前缀
		// 11101010
		// 11101111
		lo, hi = lo>>1, hi>>1
		shift++
	}
	return lo << shift // TODO
}
