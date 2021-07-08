package main

// Link: https://leetcode-cn.com/problems/excel-sheet-column-title

func convertToTitle(columnNumber int) string {
	var res []byte
	for 0 < columnNumber {
		columnNumber -= 1
		res = append(res, byte(columnNumber%26+'A'))
		columnNumber /= 26
	}

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}

	return string(res)
}
