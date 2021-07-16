package main

// https://leetcode-cn.com/problems/stone-game

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}

	for r := 0; r < n; r++ {
		for l := 0; l < r; l++ {
			dp[l][r] = max(piles[l]-dp[l+1][r], piles[r]-dp[l][r-1])
		}
	}

	return 0 < dp[0][n-1]
}

// 压缩
func stoneGame(piles []int) bool {
	top := len(piles) - 1
	dp := make([]int, top+1)
	dp[0] = piles[0]

	for r := 0; r < top; r++ {
		for l := 0; l < r; l++ {
			dp[l] = max(piles[l]-dp[l+1], piles[r]-dp[r-1])
		}
	}
	return 0 < max(piles[0]-dp[1], piles[top]-dp[top-1])
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
