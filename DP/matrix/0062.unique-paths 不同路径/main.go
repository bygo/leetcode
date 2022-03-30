package main

// https://leetcode-cn.com/problems/unique-paths

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)
func uniquePaths(row, col int) int {
	f := make([][]int, row)
	for i := range f {
		f[i] = make([]int, col)
		f[i][0] = 1
	}
	for j := 0; j < col; j++ {
		f[0][j] = 1
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[row-1][col-1]
}
