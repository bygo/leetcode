package main

// https://leetcode-cn.com/problems/hamming-distance

// ❓ 汉明距离

func hammingDistance(x int, y int) int {
	// 异或取不同个数
	z := x ^ y
	var ones int
	for 0 < z {
		// bk
		z &= z - 1
		ones++
	}
	return ones
}
