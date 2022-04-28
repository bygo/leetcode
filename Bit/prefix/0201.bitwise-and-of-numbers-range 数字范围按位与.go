package main

// https://leetcode-cn.com/problems/bitwise-and-of-numbers-range

// ❓ 数字范围按位与

func rangeBitwiseAnd(lo int, hi int) int {
	shift := 0
	for lo < hi {
		// 找出公共前缀
		lo, hi = lo>>1, hi>>1
		shift++
	}
	// 还原值
	return lo << shift
}
