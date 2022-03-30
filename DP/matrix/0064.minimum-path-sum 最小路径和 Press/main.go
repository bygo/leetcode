package main

// https://leetcode-cn.com/problems/minimum-path-sum

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)
// 压缩
func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	f := make([]int, col)
	f[0] = grid[0][0]
	for j := 1; j < col; j++ {
		f[j] = grid[0][j] + f[j-1]
	}

	for i := 1; i < row; i++ {
		f[0] += grid[i][0]
		for j := 1; j < col; j++ {
			f[j] = min(f[j-1], f[j]) + grid[i][j]
		}
	}

	return f[col-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
