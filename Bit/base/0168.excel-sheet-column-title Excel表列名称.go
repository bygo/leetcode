package main

// https://leetcode.cn/problems/excel-sheet-column-title

// ❓ Excel 列名称

func convertToTitle(num int) string {
	var lo = 7
	var buf = make([]byte, lo)
	for 0 < num {
		num--
		lo--
		buf[lo] = byte(num%26 + 'A')
		num /= 26
	}
	return string(buf[lo:])
}
