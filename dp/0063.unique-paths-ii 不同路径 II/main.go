package main

// Link: https://leetcode-cn.com/problems/unique-paths-ii

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	l1 := len(obstacleGrid)
	l2 := len(obstacleGrid[0])
	f := make([][]int, l1)

	var i int
	for i < l1 { // 第一列
		if obstacleGrid[i][0] == 1 {
			break
		}
		f[i] = make([]int, l2)
		f[i][0] = 1
		i++
	}

	for i < l1 {
		f[i] = make([]int, l2)
		i++
	}

	for j := 0; j < l2; j++ { // 第一行
		if obstacleGrid[0][j] == 1 {
			break
		}
		f[0][j] = 1
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if obstacleGrid[i][j] == 0 {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[l1-1][l2-1]
}

// 压缩
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	l1, l2 := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, l2)
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	f[0] = 1

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
			} else if 1 <= j {
				f[j] += f[j-1]
			}
		}
	}
	return f[l2-1]
}
