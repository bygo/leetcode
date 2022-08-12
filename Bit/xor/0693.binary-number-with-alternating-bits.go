package main

// https://leetcode.cn/problems/binary-number-with-alternating-bits

// ❓ 交替位二进制数

func hasAlternatingBits(n int) bool {
	// 全部变1
	n = n ^ (n >> 1) + 1
	// 全部进位成功 只剩1个1
	return n&(n-1) == 0
}
