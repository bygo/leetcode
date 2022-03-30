package main

// https://leetcode-cn.com/problems/unique-paths-ii

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)

// 压缩
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, col)
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	f[0] = 1

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
			} else if 1 <= j {
				f[j] += f[j-1]
			}
		}
	}
	return f[col-1]
}
