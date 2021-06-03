package main

//Link: https://leetcode-cn.com/problems/unique-paths-ii

// dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)

	var i int
	for i < m { // 第一列
		if obstacleGrid[i][0] == 1 {
			for i < m {
				dp[i] = make([]int, n)
				i++
			}
		} else {
			dp[i] = make([]int, n)
			dp[i][0] = 1
			i++
		}
	}

	for j := 0; j < n; j++ { // 第一行
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 压缩
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, m)
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if 1 <= j {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[m-1]
}
