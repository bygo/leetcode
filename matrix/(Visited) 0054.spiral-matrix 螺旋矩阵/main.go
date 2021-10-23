package main

// https://leetcode-cn.com/problems/spiral-matrix

func spiralOrder(matrix [][]int) []int {
	var (
		rows    = len(matrix)
		cols    = len(matrix[0])
		total   = rows * cols
		visited = make([][]bool, rows)
	)

	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var (
		res      = make([]int, total)
		dirs     = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		row, col int
		cur      int
	)
	for i := 0; i < total; i++ {
		res[i] = matrix[row][col]
		visited[row][col] = true
		row += dirs[cur][0]
		col += dirs[cur][1]
		if row < 0 || rows <= row || col < 0 || cols <= col || visited[row][col] {
			row -= dirs[cur][0]
			col -= dirs[cur][1]
			cur = (cur + 1) % 4
			row += dirs[cur][0]
			col += dirs[cur][1]
		}
	}
	return res
}
