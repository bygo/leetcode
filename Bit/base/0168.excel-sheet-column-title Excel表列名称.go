package main

// https://leetcode-cn.com/problems/excel-sheet-column-title

// ❓ Excel 列名称

func convertToTitle(num int) string {
	var buf = make([]byte, 7)
	var lo = 7
	for 0 < num {
		lo--
		num -= 1
		buf[lo] = byte(num%26) + 'A'
		num /= 26
	}
	return string(buf[lo:])
}
