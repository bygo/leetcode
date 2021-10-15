package main

// https://leetcode-cn.com/problems/sum-of-two-integers

func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1 // 进位
		a ^= b                  // 求不进位和
		b = int(carry)
	}
	return a
}
