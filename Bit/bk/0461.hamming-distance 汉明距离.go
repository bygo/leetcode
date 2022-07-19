package main

// https://leetcode.cn/problems/hamming-distance

// ❓ 汉明距离

func hammingDistance(x int, y int) int {
	// xor 取不同个数
	num := x ^ y
	// bk
	var cnt int
	for 0 < num {
		// bk
		num &= num - 1
		cnt++
	}
	return cnt
}
