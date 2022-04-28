package main

// https://leetcode-cn.com/problems/arranging-coins

func arrangeCoins(n int) int {
	lo, hi := 0, n+1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		cur := mid * (mid + 1) / 2
		if cur <= n {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}
