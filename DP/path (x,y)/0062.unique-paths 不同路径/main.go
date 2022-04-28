package main

// https://leetcode-cn.com/problems/unique-paths

// f(i)(j) = f(i-1)(j) + f(i)(j-1)

func uniquePaths(row, col int) int {
	f := make([][]int, row)
	// 第一行
	for i := range f {
		f[i] = make([]int, col)
		f[i][0] = 1
	}
	// 第一列
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

// 压缩
func uniquePaths_(row, col int) int {
	f := make([]int, col)
	f[0] = 1
	for i := 0; i < row; i++ {
		for j := 1; j < col; j++ {
			f[j] += f[j-1]
		}
	}
	return f[col-1]
}
