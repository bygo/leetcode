package main

// https://leetcode.cn/problems/longest-increasing-path-in-a-matrix

func longestIncreasingPath(matrix [][]int) int {
	var (
		dirs       = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		rows, cols = len(matrix), len(matrix[0])
	)
	degree := make([][]int, rows)

	for i := 0; i < rows; i++ {
		degree[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range dirs {
				nr, nc := i+dir[0], j+dir[1]
				if 0 <= nr && nr < rows && 0 <= nc && nc < cols && matrix[i][j] < matrix[nr][nc] {
					degree[i][j]++ // 入度
				}
			}
		}
	}

	// 取出所有0度
	queue := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if degree[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	res := 0
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		res++
		for i := 0; i < cnt; i++ {
			q := queue[i]
			row, col := q[0], q[1]
			for _, dir := range dirs {
				nr, nc := row+dir[0], col+dir[1]
				if 0 <= nr && nr < rows && 0 <= nc && nc < cols && matrix[nr][nc] < matrix[row][col] {
					degree[nr][nc]-- // 出度
					if degree[nr][nc] == 0 {
						queue = append(queue, []int{nr, nc})
					}
				}
			}
		}
		queue = queue[cnt:]
	}
	return res
}
