package main

//Link: https://leetcode-cn.com/problems/guess-number-higher-or-lower

func guessNumber(r int) int {
	var l = 1
	for l <= r {
		mid := l + (r-l)/2
		switch guess(mid) {
		case -1:
			r = mid - 1
		case 1:
			l = mid + 1
		case 0:
			return mid
		}
	}
	return -1
}

func guess(n int) int {
	return 0
}
