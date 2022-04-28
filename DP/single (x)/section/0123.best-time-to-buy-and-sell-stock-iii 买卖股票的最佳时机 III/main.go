package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii

//
func maxProfit(prices []int) int {
	pL := len(prices)
	b1, b2 := -prices[0], -prices[0]
	s1, s2 := 0, 0
	for day := 1; day < pL; day++ {
		// 买入最优解
		b1 = max(b1, -prices[day])
		// 买入最优解 + 卖出最优解
		s1 = max(s1, b1+prices[day])

		// 前置最优解 + 买入最优解
		b2 = max(b2, s1-prices[day])
		// 买入最优解 + 卖出最优解
		s2 = max(s2, b2+prices[day])
	}

	return s2
}

func maxProfit_(prices []int) int {
	// 卖出次数
	pL := len(prices)
	f := make([][3][2]int, pL)

	for cnt := 1; cnt <= 2; cnt++ {
		f[0][cnt][0] = 0
		f[0][cnt][1] = -prices[0]
		for day := 1; day < pL; day++ {
			// 买入 = 前一天最优解 |  前一次卖出最优解 - 价格
			f[day][cnt][1] = max(f[day-1][cnt][1], f[day-1][cnt-1][0]-prices[day])
			// 卖出 = 前一天最优解 |  前一次买入最优解 + 价格
			f[day][cnt][0] = max(f[day-1][cnt][0], f[day-1][cnt][1]+prices[day])
		}
	}

	return f[pL-1][2][0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
