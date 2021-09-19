package main

// https://leetcode-cn.com/problems/valid-perfect-square

func isPerfectSquare(num int) bool {
	l, r := 0, num+1
	for l < r {
		mid := l + (r-l)/2
		cur := mid * mid
		if cur == num {
			return true
		} else if cur < num {
			l = mid + 1
		} else if num < cur {
			r = mid
		}
	}
	return false
}
