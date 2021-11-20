package main

// https://leetcode-cn.com/problems/hamming-distance

func hammingDistance(x int, y int) int {
	z := x ^ y
	var res int
	for 0 < z {
		z &= z - 1
		res++
	}
	return res
}
