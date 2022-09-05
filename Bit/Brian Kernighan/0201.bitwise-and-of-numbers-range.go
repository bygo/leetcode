package main

// https://leetcode.cn/problems/bitwise-and-of-numbers-range

func rangeBitwiseAnd(lo int, hi int) int {
	for lo < hi { // TODO 按位与 快速缩进
		hi &= hi - 1
	}
	return hi
}
