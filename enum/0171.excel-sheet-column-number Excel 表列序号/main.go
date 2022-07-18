package main

// https://leetcode.cn/problems/excel-sheet-column-number

func titleToNumber(columnTitle string) int {
	var res int
	for i := range columnTitle {
		res = res*26 + int(columnTitle[i]-'@')
	}
	return res
}
