package main

// https://leetcode-cn.com/problems/bitwise-and-of-numbers-range

// ❓ 数字范围按位与

func rangeBitwiseAnd(lo int, hi int) int {
	for lo < hi {
		// 每次消除与下一个数字的 1，数字会快速变小
		hi &= hi - 1
	}
	return hi
}
