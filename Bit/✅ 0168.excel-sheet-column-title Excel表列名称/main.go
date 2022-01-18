package main

// https://leetcode-cn.com/problems/excel-sheet-column-title

// ❓ Excel 列名称

func convertToTitle(columnNumber int) string {
	var title = []byte{}
	for 0 < columnNumber {
		// 偏移 A从1开始
		columnNumber -= 1
		title = append(title, byte(columnNumber%26+'A'))
		columnNumber /= 26
	}
	l, r := 0, len(title)-1
	for l < r {
		title[l], title[r] = title[r], title[l]
		l++
		r--
	}
	return string(title)
}
