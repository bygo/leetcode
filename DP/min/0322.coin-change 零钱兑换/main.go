package main

// https://leetcode.cn/problems/coin-change

// f[i] = f[i-coin] + 1
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)

	f[0] = 0
	for i := 1; i <= amount; i++ {
		f[i] = amount + 1
		for _, coin := range coins {
			if coin <= i {
				f[i] = min(f[i], f[i-coin]+1)
			}
		}
	}
	if amount < f[amount] {
		return -1
	}
	return f[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
