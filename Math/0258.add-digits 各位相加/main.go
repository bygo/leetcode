package main

// https://leetcode.cn/problems/add-digits

// ❓ 各位相加
// ⚠️ 数根

func addDigits(num int) int {
	return (num-1)%9 + 1
}
