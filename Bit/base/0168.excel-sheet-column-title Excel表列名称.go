package main

// https://leetcode-cn.com/problems/excel-sheet-column-title

// ❓ Excel 列名称

func convertToTitle(num int) string {
	var buf []byte

	for 0 < num {
		// 偏移 A从1开始
		// 0~25 -> 1~26
		// 也就是 26(z)不进位，27(aa)才进位
		num -= 1
		buf = append(buf, byte(num%26)+'A')
		num /= 26
	}

	lo, hi := 0, len(buf)-1
	for lo < hi {
		buf[lo], buf[hi] = buf[hi], buf[lo]
		lo++
		hi--
	}
	return string(buf)
}

func convertToTitle(num int) string {
	var buf []byte
	for 0 < num {
		num -= 1
		buf = append(buf, byte(num%26)+'A')
		num /= 26
	}
}
