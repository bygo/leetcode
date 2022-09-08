package main

// https://leetcode.cn/problems/bitwise-and-of-numbers-range

// 范围与
func rangeBitwiseAnd(lo int, hi int) int {
	for lo < hi { // TODO
		hi &= hi - 1
	}
	return hi
}
