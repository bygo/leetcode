package main

// https://leetcode.cn/problems/excel-sheet-column-title

func convertToTitle(num int) string {
	var lo = 7
	var buf = make([]byte, lo) // TODO 避免数组反转
	for 0 < num {
		num-- // TODO 修正偏移
		lo--
		buf[lo] = byte(num%26 + 'A')
		num /= 26
	}
	return string(buf[lo:])
}
