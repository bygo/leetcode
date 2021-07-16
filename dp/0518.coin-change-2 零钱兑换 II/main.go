package main

// https://leetcode-cn.com/problems/coin-change-2

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	//多加一种硬币 ，dp[i] 可以多几种跳跃方式
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if coin <= i {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[amount]
}

func wrong(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	// 多加一块钱 dp[i] 可以多几种跳跃方式  like（斐波那契数）
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[amount]
}
