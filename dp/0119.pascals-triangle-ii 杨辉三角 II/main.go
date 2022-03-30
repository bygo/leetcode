package main

// https://leetcode-cn.com/problems/pascals-triangle-ii

// 压缩
// f(i)(j) = f(i-1)(j-1) + f(i-1)(j)
func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	var f = make([]int, numRows)
	for i := 0; i < numRows; i++ {
		f[0] = 1
		f[i] = 1
		for j := i - 1; 0 < j; j-- {
			f[j] = f[j-1] + f[j]
		}
	}
	return f
}
