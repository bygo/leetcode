package main

func maxProfit(prices []int) int {
	cur := prices[0]
	res := 0
	l := len(prices)
	for i := 1; i < l; i++ {
		if cur < prices[i] {
			res += prices[i] - cur
		}
		cur = prices[i]
	}
	return res
}
