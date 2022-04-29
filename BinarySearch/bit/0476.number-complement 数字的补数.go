package main

// https://leetcode-cn.com/problems/number-complement

// ❓ 数字的补数
// ⚠️ 补数 = 最大值 - num

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
	// 补数 = 1<<mid - 1 - num
	return 1<<hi - 1 - num
}
