package main

// https://leetcode-cn.com/problems/sqrtx

func mySqrt(x int) int {
	var l = 0
	var r = x
	for l < r {
		mid := l + (r-l)/2
		cur := mid * mid
		if cur < x {
			l = mid + 1
		} else if x < cur {
			r = mid
		} else if x == cur {
			return mid
		}
	}
	return l - 1
}
