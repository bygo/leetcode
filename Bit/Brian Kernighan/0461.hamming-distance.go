package main

// https://leetcode.cn/problems/hamming-distance

func hammingDistance(x int, y int) int {
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
