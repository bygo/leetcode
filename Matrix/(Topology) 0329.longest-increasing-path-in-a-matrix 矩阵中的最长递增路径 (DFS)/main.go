package main

// https://leetcode.cn/problems/longest-increasing-path-in-a-matrix

func longestIncreasingPath(matrix [][]int) int {
	var (
		dirs       = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		rows, cols = len(matrix), len(matrix[0])
	)

	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
	}
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if memo[row][col] != 0 {
			return memo[row][col]
		}
		memo[row][col]++
		for _, dir := range dirs {
			nr, nc := row+dir[0], col+dir[1]
			if 0 <= nr && nr < rows && 0 <= nc && nc < cols && matrix[row][col] < matrix[nr][nc] {
				memo[row][col] = max(memo[row][col], dfs(nr, nc)+1)
			}
		}
		return memo[row][col]
	}
	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res = max(res, dfs(i, j))
		}
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
