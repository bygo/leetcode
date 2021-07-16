package main

// https://leetcode-cn.com/problems/minimum-path-sum

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)
func minPathSum(grid [][]int) int {
	l1, l2 := len(grid), len(grid[0])
	f := make([][]int, l1)
	f[0] = make([]int, l2)
	f[0][0] = grid[0][0]
	for i := 1; i < l1; i++ {
		f[i] = make([]int, l2)
		f[i][0] = grid[i][0] + f[i-1][0]
	}

	for i := 1; i < l2; i++ {
		f[0][i] = grid[0][i] + f[0][i-1]
	}

	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
		}
	}

	return f[l1-1][l2-1]
}

// 压缩
func minPathSum(grid [][]int) int {
	l1, l2 := len(grid), len(grid[0])
	f := make([]int, l2)
	f[0] = grid[0][0]
	for j := 1; j < l2; j++ {
		f[j] = grid[0][j] + f[j-1]
	}

	for i := 1; i < l1; i++ {
		f[0] += grid[i][0]
		for j := 1; j < l2; j++ {
			f[j] = min(f[j-1], f[j]) + grid[i][j]
		}
	}

	return f[l2-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
