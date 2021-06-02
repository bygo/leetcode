package main

//Link: https://leetcode-cn.com/problems/unique-paths-ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, m)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	} else {
		return 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if 1 <= j {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[len(dp)-1]
}

//func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//	m := len(obstacleGrid)
//	n := len(obstacleGrid[0])
//	dp := make([][]int, m)
//
//	stop := false
//	for i := range dp { // 第一列
//		dp[i] = make([]int, n)
//		if obstacleGrid[i][0] == 1 {
//			stop = true
//		}
//		if !stop {
//			dp[i][0] = 1
//		}
//	}
//
//	for j := 0; j < n; j++ { // 第一行
//		if obstacleGrid[0][j] == 1 {
//			break
//		}
//		dp[0][j] = 1
//	}
//	for i := 1; i < m; i++ {
//		for j := 1; j < n; j++ {
//			if obstacleGrid[i][j] == 1 {
//				dp[i][j] = 0
//			} else {
//				dp[i][j] = dp[i-1][j] + dp[i][j-1]
//			}
//		}
//	}
//	return dp[m-1][n-1]
//}
