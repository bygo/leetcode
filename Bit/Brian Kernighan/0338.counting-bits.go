package main

// https://leetcode.cn/problems/counting-bits

func countBits(numMax int) []int {
	var bits []int
	for num := 0; num <= numMax; num++ {
		bits = append(bits, bk(num))
	}
	return bits
}

func bk(num int) int {
	var cnt int
	for 0 < num {
		num &= num - 1
		cnt++
	}
	return cnt
}
