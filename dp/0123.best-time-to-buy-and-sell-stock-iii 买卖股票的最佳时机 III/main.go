package main

// Link: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii

func maxProfit(prices []int) int {
	a, b := -prices[0], 0
	c, d := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		a = max(a, -prices[i])  // 最少钱购买
		b = max(b, a+prices[i]) // 最多钱卖出
		c = max(c, b-prices[i]) // 最少钱购买
		d = max(d, c+prices[i]) // 最多钱卖出
	}
	return d
}

func maxProfit(prices []int) int {
	count := 2
	n := len(prices)
	dp := make([][3][2]int, n)

	for i := 0; i < len(prices); i++ {
		for j := count; 1 <= j; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}

			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j-1][0]-prices[i], dp[i-1][j][1])
		}
	}

	return dp[len(prices)-1][count][0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
