package main

// https://leetcode-cn.com/problems/count-operations-to-obtain-zero

// ❓ 得到 0 的操作数

func countOperations(num1, num2 int) int {
	var cnt int
	for 0 < num1 {
		cnt += num2 / num1
		num1, num2 = num2%num1, num1
	}
	return cnt
}
