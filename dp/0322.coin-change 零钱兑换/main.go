package main

//Link: https://leetcode-cn.com/problems/coin-change

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 0
	// 每种amount 有多少种组合
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if amount < dp[amount] {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1) // 1 1  0 2
			}
		}
	}

	if amount < dp[amount] {
		return -1
	}

	return dp[amount]
}
