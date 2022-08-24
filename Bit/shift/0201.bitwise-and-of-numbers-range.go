package main

// https://leetcode.cn/problems/bitwise-and-of-numbers-range

func rangeBitwiseAnd(lo int, hi int) int {
	for lo < hi {
		hi &= hi - 1
	}
	return hi
}

func rangeBitwiseAnd(lo int, hi int) int {
	shift := 0
	for lo < hi {
		lo, hi = lo>>1, hi>>1
		shift++
	}
	return lo << shift
}
