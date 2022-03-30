package main

// https://leetcode-cn.com/problems/unique-paths-ii

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	f := make([][]int, row)

	// 第一列
	var idx int
	for idx < row {
		if obstacleGrid[idx][0] == 1 {
			break
		}
		f[idx] = make([]int, col)
		f[idx][0] = 1
		idx++
	}
	for idx < row {
		f[idx] = make([]int, col)
		idx++
	}

	// 第一行
	for j := 0; j < col; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		f[0][j] = 1
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 0 {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[row-1][col-1]
}
