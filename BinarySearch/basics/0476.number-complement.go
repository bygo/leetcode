package main

// https://leetcode.cn/problems/number-complement

func findComplement(num int) int {
	lo, hi := 0, 31
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		numCur := 1 << mid
		if numCur <= num {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return 1<<hi - 1 - num
}
