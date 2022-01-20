package main

// https://leetcode-cn.com/problems/excel-sheet-column-title

// ❓ Excel 列名称

func convertToTitle(columnNumber int) string {
	var titleBuf = []byte{}
	for 0 < columnNumber {
		// 偏移 A从1开始
		// 0~26 -> 1~27
		// 也就是 26(z)不进位，27(aa)才进位
		columnNumber -= 1
		titleBuf = append(titleBuf, byte((columnNumber)%26+'A'))
		columnNumber /= 26
	}

	lo, hi := 0, len(titleBuf)-1
	for lo < hi {
		titleBuf[lo], titleBuf[hi] = titleBuf[hi], titleBuf[lo]
		lo++
		hi--
	}
	return string(titleBuf)
}
