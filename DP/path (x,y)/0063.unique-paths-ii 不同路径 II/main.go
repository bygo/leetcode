package main

// https://leetcode-cn.com/problems/unique-paths-ii

// f(i)(j) = f(i-1)(j) + f(i)(j-1)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	f := make([][]int, row)

	// 第一列
	var idx int
	for idx < row && obstacleGrid[idx][0] == 0 {
		// 阻塞0
		f[idx] = make([]int, col)
		f[idx][0] = 1
		idx++
	}
	// 第一列 补充
	for idx < row {
		f[idx] = make([]int, col)
		idx++
	}

	// 第一行
	for j := 0; j < col && obstacleGrid[0][j] == 0; j++ {
		// 阻塞0
		f[0][j] = 1
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 0 {
				// 阻塞0
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[row-1][col-1]
}

// 压缩
func uniquePathsWithObstacles_(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, col)
	f[0] = 1
	for i := 0; i < row; i++ {
		if obstacleGrid[i][0] == 1 {
			f[0] = 0
		}
		for j := 1; j < col; j++ {
			// 阻塞为0
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
			} else {
				f[j] = f[j] + f[j-1]
			}
		}
	}
	return f[col-1]
}
