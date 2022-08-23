package main

// https://leetcode.cn/problems/convert-integer-lcci/

// ❓ 整数转换

func convertInteger(a int, b int) int {
	// xor
	n := int32(a) ^ int32(b)
	// bk
	var cnt int
	for 0 != n {
		n &= n - 1
		cnt++
	}
	return cnt
}
