package main

func maxProfit(prices []int) int {
	cur := prices[0]
	earn := 0
	l := len(prices)
	for i := 1; i < l; i++ {
		if cur < prices[i] {
			earn += prices[i] - cur
		}
		cur = prices[i]
	}
	return earn
}
