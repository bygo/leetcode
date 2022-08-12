package main

// https://leetcode.cn/problems/guess-number-higher-or-lower

func guessNumber(hi int) int {
	var lo = 1
	hi += 1
	for {
		mid := int(uint(lo+hi) >> 1)
		cur := guess(mid)
		if cur == 1 {
			lo = mid + 1
		} else if cur == -1 {
			hi = mid
		} else if cur == 0 {
			return mid
		}
	}
}

func guess(n int) int {
	return 0 // -1 0 1
}
