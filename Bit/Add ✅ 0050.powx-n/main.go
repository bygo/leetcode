package main

// https://leetcode-cn.com/problems/powx-n

// ❓ 实现 Pow(x, n)

func myPow(x float64, n int) float64 {
	var res float64 = 1
	var num = n
	if num < 0 {
		num = -num
	}

	for 0 < num {
		// 快速乘
		if 1 == num&1 {
			res *= x
		}
		x *= x // x<<=1
		num >>= 1
	}
	// 负数倒数
	if n < 0 {
		return 1 / res
	}
	return res
}
