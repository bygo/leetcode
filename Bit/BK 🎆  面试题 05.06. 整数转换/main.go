package main

// https://leetcode-cn.com/problems/convert-integer-lcci/

// ❓ 整数转换

// ⚠️ 符号位也会参与运算

func convertInteger(a int, b int) int {
	// xor
	n := int32(a) ^ int32(b)
	// bk
	var cnt int
	for 0 != n { // ⚠️ 正->负数
		n &= n - 1 // 负数依然正常运算,当负数超过1000，溢出为0111
		cnt++
	}
	return cnt
}
