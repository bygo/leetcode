package main

// https://leetcode-cn.com/problems/counting-bits

// ❓ 比特位计数
// ⚠️ 0~n 有几个1

func countBits(n int) []int {
	var bits []int
	for num := 0; num <= n; num++ {
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
