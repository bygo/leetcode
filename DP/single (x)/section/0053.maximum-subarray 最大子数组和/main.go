package main

// https://leetcode-cn.com/problems/maximum-subarray

// 前缀
// f(n) = f(n) + f(n-1)
func maxSubArray(f []int) int {
	sumMax := f[0]
	fL := len(f)
	for i := 1; i < fL; i++ {
		if 0 < f[i-1] { // 有所贡献
			f[i] += f[i-1] // 保证子数组最大
		}
		if sumMax < f[i] {
			sumMax = f[i]
		}
	}
	return sumMax
}
