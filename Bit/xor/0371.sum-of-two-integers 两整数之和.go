package main

// https://leetcode-cn.com/problems/sum-of-two-integers

// ❓ 两整数之和
// ⚠️ 不使用+ - 符号
// ⚠️ 有负数

func getSum(a int, b int) int {
	for 0 != b { // 过大的数，可通过转换为 uint ,防止^负数溢出 （变正）
		tmp := a & b << 1 // 进位，直至 `b` 为 0，不再需要进位,
		a = a ^ b         // 不进位，a ^ b 已加上 `b` 剩下值
		b = tmp           // 进位数
	}
	return a
}
