package main

// https://leetcode-cn.com/problems/powx-n

// ❓ 实现 Pow(x, n)

func myPow(num float64, pow int) float64 {
	var numRes float64 = 1
	// 负指数幂
	var minus bool
	if pow < 0 {
		minus = true
		pow = -pow
	}

	for 0 < pow {
		// 快速乘
		if pow&1 == 1 {
			numRes *= num
		}
		num *= num
		pow >>= 1
	}

	// 负数倒数
	if minus {
		return 1 / numRes
	}
	return numRes
}
