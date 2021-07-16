package main

// https://leetcode-cn.com/problems/maximum-subarray

// 前缀
// f(n) = f(n) + f(n-1)
func maxSubArray(f []int) int {
	res := f[0]
	l1 := len(f)
	for i := 1; i < l1; i++ {
		if 0 < f[i-1] { // 贡献
			f[i] += f[i-1]
		}
		if res < f[i] {
			res = f[i]
		}
	}
	return res
}
