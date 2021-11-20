package main

// https://leetcode-cn.com/problems/powx-n

func myPow(x float64, n int) float64 {
	var res float64 = 1
	var y = n
	if y < 0 {
		y = -y
	}
	for 0 < y {
		if 1 == y&1 {
			res *= x
		}
		x *= x
		y >>= 1
	}
	if n < 0 {
		return 1 / res
	}
	return res
}
