package main

// https://leetcode-cn.com/problems/valid-perfect-square

func isPerfectSquare(num int) bool {
	lo, hi := 0, num+1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		pow := mid * mid
		if pow < num {
			lo = mid + 1
		} else if num < pow {
			hi = mid
		} else if pow == num {
			return true
		}
	}
	return false
}
