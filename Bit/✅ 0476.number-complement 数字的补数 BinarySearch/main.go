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

	// num 补全数 1<<mid - 1
	// num 补数 = 补全数 - num
	return 1<<lo - 1 - num
}
