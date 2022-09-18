package main

// https://leetcode.cn/problems/complement-of-base-10-integer

func bitwiseComplement(n int) int {
	lo, hi := 1, 31 // TODO
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		numCur := 1 << mid
		if numCur <= n {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return 1<<hi - 1 ^ n
}
