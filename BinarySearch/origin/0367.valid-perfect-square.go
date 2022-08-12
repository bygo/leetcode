package main

// https://leetcode.cn/problems/valid-perfect-square

func isPerfectSquare(num int) bool {
	lo, hi := 0, num+1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		pow := mid * mid
		if pow < num {
			lo = mid + 1
		} else if num < pow {
			hi = mid
		} else if num == pow {
			return true
		}
	}
	return false
}

func isPerfectSquare(num int) bool {
	lo, hi := 0, num+1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		pow := mid * mid
		if pow < num {
			lo = mid + 1
		} else if num < pow {
			hi = mid
		} else if num == pow {
			return true
		}
	}
	return false
}
