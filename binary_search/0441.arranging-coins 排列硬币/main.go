package main

// https://leetcode-cn.com/problems/arranging-coins

func arrangeCoins(n int) int {
	l, r := 0, n+1
	for l < r {
		mid := l + (r-l)/2
		cur := mid * (mid + 1) / 2
		if cur <= n {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}
