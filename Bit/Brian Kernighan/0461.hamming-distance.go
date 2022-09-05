package main

// https://leetcode.cn/problems/hamming-distance

func hammingDistance(x int, y int) int {
	num := x ^ y // TODO 两数 位差
	var cnt int
	for 0 < num {
		num &= num - 1
		cnt++
	}
	return cnt
}
