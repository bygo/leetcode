package main

// https://leetcode.cn/problems/arranging-coins

// general
func arrangeCoins(n int) int {
	lo, hi := 0, n+1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		cnt := mid * (mid + 1) / 2 // TODO 等差数列
		if cnt < n {
			lo = mid + 1
		} else if n < cnt {
			hi = mid
		} else if cnt == n {
			return mid
		}
	}
	return lo - 1
}

// offset
func arrangeCoins(n int) int {
	lo, hi := -1, n
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		cnt := mid * (mid + 1) / 2
		if cnt < n {
			lo = mid
		} else if n < cnt {
			hi = mid - 1
		} else if n == cnt {
			return mid
		}
	}
	return lo
}

// equal
func arrangeCoins(n int) int {
	lo, hi := 0, n
	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)
		cnt := mid * (mid + 1) / 2
		if cnt < n {
			lo = mid + 1
		} else if n < cnt {
			hi = mid - 1
		} else if n == cnt {
			return mid
		}
	}
	return hi
}
