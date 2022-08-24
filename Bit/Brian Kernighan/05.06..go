package main

// https://leetcode.cn/problems/convert-integer-lcci/

func convertInteger(a int, b int) int {
	n := int32(a) ^ int32(b) // TODO
	var cnt int
	for 0 != n {
		n &= n - 1 // TODO
		cnt++
	}
	return cnt
}
