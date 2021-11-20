package main

// https://leetcode-cn.com/problems/number-complement

func findComplement(num int) int {
	i := 1
	l, r := 0, 31
	for l < r {
		mid := l + (r-l)/2
		if i<<mid <= num {
			l = mid + 1
		} else {
			r = mid
		}
	}
	i <<= l
	return i - num - 1
}
