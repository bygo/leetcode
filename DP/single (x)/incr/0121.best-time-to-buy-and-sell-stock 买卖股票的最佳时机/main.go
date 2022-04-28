package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

func maxProfit(prices []int) int {
	var minBuy, maxSale = 1<<63 - 1, 0
	for _, price := range prices {
		if price < minBuy {
			minBuy = price
		}
		if maxSale < price-minBuy {
			maxSale = price - minBuy
		}
	}
	return maxSale
}
