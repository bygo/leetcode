package main

// https://leetcode.cn/problems/pascals-triangle

// 二维
// f(i)(j) = f(i-1)(j-1) + f(i-1)(j)
func generate(numRows int) [][]int {
	var f = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		f[i] = make([]int, i+1)
		f[i][0] = 1
		f[i][i] = 1
		for j := 1; j < i; j++ {
			f[i][j] = f[i-1][j-1] + f[i-1][j]
		}
	}

	return f
}
