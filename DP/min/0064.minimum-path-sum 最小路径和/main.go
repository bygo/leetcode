package main

// https://leetcode-cn.com/problems/minimum-path-sum

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)
// 📚 最优路径选择

// 压缩
func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	f := make([]int, col)
	f[0] = grid[0][0]
	for j := 1; j < col; j++ {
		f[j] = f[j-1] + grid[0][j]
	}

	for i := 1; i < row; i++ {
		f[0] += grid[i][0]
		for j := 1; j < col; j++ {
			f[j] = min(f[j-1], f[j]) + grid[i][j]
		}
	}

	return f[col-1]
}

func _(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	f := make([][]int, row)
	f[0] = make([]int, col)
	f[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		f[i] = make([]int, col)
		f[i][0] = grid[i][0] + f[i-1][0]
	}

	for i := 1; i < col; i++ {
		f[0][i] = grid[0][i] + f[0][i-1]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
		}
	}

	return f[row-1][col-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
