package main

// https://leetcode-cn.com/problems/unique-paths

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)

// 压缩
func uniquePaths(row, col int) int {
	f := make([]int, col)
	f[0] = 1
	for i := 0; i < row; i++ {
		for j := 1; j < col; j++ {
			f[j] += f[j-1]
		}
	}
	return f[col-1]
}
