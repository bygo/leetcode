package main

// https://leetcode-cn.com/problems/guess-number-higher-or-lower

func guessNumber(r int) int {
	var l = 1
	for l <= r {
		mid := l + (r-l)/2
		cur := guess(mid)
		if cur == -1 {
			r = mid - 1
		} else if cur == 1 {
			l = mid + 1
		} else if cur == 0 {
			return mid
		}
	}
	return -1
}

func guess(n int) int {
	return 0
}
