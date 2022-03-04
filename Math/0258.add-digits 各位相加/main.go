package main

// https://leetcode-cn.com/problems/add-digits

// ❓ 各位相加

func addDigits(num int) int {
	return (num-1)%9 + 1
}
