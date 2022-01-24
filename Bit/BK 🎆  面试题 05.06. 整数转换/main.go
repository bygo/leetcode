package main

// https://leetcode-cn.com/problems/convert-integer-lcci/

// ❓ 整数转换

// ⚠️ 符号位也会参与运算

func convertInteger(a int, b int) int {
	n := int32(a) ^ int32(b)
	var cnt int
	for 0 != n { // ⚠️ 负数
		n &= n - 1
		cnt++
	}
	return cnt
}
