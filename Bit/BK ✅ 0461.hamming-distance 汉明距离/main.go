package main

// https://leetcode-cn.com/problems/hamming-distance

// ❓ 汉明距离

func hammingDistance(x int, y int) int {
	// 异或取不同个数
	num := x ^ y
	var cntDiff int
	for 0 < num {
		// bk
		num &= num - 1
		cntDiff++
	}
	return cntDiff
}
