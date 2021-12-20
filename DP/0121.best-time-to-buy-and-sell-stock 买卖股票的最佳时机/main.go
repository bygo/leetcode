package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

func maxProfit(prices []int) int {
	var min, max = 1<<63 - 1, 0
	for _, p := range prices {
		if p < min {
			min = p
		}
		if max < p-min {
			max = p - min
		}
	}
	return max
}
