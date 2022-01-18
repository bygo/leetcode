package main

// https://leetcode-cn.com/problems/sum-of-two-integers

// ❓ 两整数之和
// ⚠️ 不使用+ - 符号

func getSum(a int, b int) int {
	for b != 0 { // 通过转换为 uint ,防止负数溢出 （变正）
		carry := uint(a&b) << 1 // 进位，直至 b 为 0，不再需要进位,
		a = a ^ b               // 不进位时，a ^ b 已经加上b剩下所有值
		b = int(carry)
	}
	return a
}
