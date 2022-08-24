package main

// https://leetcode.cn/problems/hamming-distance

func hammingDistance(x int, y int) int {
	num := x ^ y
	var cnt int
	for 0 < num {
		num &= num - 1
		cnt++
	}
	return cnt
}
