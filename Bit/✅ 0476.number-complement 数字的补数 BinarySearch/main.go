package main

// https://leetcode-cn.com/problems/number-complement

// ❓ 数字的补数

func findComplement(num int) int {
	lo, hi := 0, 31
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if 1<<mid <= num {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return 1<<lo - num - 1
}
